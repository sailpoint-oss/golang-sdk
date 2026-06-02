#!/usr/bin/env node
/**
 * build-versioned-sdk.js
 *
 * Builds one Go SDK package per resource partition found in the apis/ directory.
 * New partitions are discovered automatically — no script updates required when
 * new endpoints are added.
 *
 * Pipeline for each partition:
 *   1. Copy apis/ to .sdk-build-tmp/  (git-ignored, so source files stay clean)
 *   2. Apply prescript YAML fixes to the temp copy
 *   3. Bundle the partition openapi.yaml with redocly (resolves shared/ $refs)
 *   4. Run openapi-generator-cli with a dynamically generated config
 *   5. Run postscript.js on the generated output
 *
 * On failure, structured error logs are written to build-errors/:
 *   build-errors/summary.md              — overview of all failures
 *   build-errors/<partition>-error.md    — self-contained per-partition report
 *                                          with error output + spec file contents
 *                                          (designed to be fed directly to an AI)
 *
 * Usage:
 *   node sdk-resources/build-versioned-sdk.js <path-to-apis-dir> [--partition <name>] [--keep-tmp]
 *
 * Options:
 *   --partition <name>   Build only the named partition (default: all)
 *   --keep-tmp           Do not delete .sdk-build-tmp after the build
 */

"use strict";

const fs   = require("fs");
const path = require("path");
const { spawnSync } = require("child_process");

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

const SDK_ROOT      = path.resolve(__dirname, "..");
const TEMP_DIR      = path.join(SDK_ROOT, ".sdk-build-tmp");
const BUNDLED_DIR   = path.join(TEMP_DIR, "bundled");
const ERROR_DIR     = path.join(SDK_ROOT, "build-errors");
const JAR           = path.join(SDK_ROOT, "openapi-generator-cli.jar");
const POSTSCRIPT    = path.join(__dirname, "postscript.js");
const TEMPLATE_DIR  = path.join(__dirname, "resources");

const GIT_REPO_ID   = "golang-sdk/v2";
const GIT_USER_ID   = "sailpoint-oss";
const API_VERSION   = "v1";
const FN_PREFIX     = "V1";

// ---------------------------------------------------------------------------
// CLI args
// ---------------------------------------------------------------------------

const args = process.argv.slice(2);
if (args.length === 0 || args[0].startsWith("--")) {
  console.error("Usage: node sdk-resources/build-versioned-sdk.js <path-to-apis-dir> [--partition <name>] [--keep-tmp]");
  process.exit(1);
}

const apisDir        = path.resolve(args[0]);
const keepTmp        = args.includes("--keep-tmp");
const partitionIdx   = args.indexOf("--partition");
const onlyPartition  = partitionIdx !== -1 ? args[partitionIdx + 1] : null;

// ---------------------------------------------------------------------------
// Utility: copy directory recursively
// ---------------------------------------------------------------------------

function copyDirSync(src, dest) {
  fs.mkdirSync(dest, { recursive: true });
  for (const entry of fs.readdirSync(src, { withFileTypes: true })) {
    const srcPath  = path.join(src, entry.name);
    const destPath = path.join(dest, entry.name);
    if (entry.isDirectory()) {
      copyDirSync(srcPath, destPath);
    } else {
      fs.copyFileSync(srcPath, destPath);
    }
  }
}

// ---------------------------------------------------------------------------
// Utility: walk directory, return all file paths
// ---------------------------------------------------------------------------

function walkSync(dir, files = []) {
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      walkSync(full, files);
    } else {
      files.push(full);
    }
  }
  return files;
}

// ---------------------------------------------------------------------------
// Utility: read a file safely (returns empty string if missing)
// ---------------------------------------------------------------------------

function readFileSafe(filePath) {
  try {
    return fs.readFileSync(filePath, "utf8");
  } catch {
    return "";
  }
}

// ---------------------------------------------------------------------------
// Prescript fixes (applied to the temp copy of apis/)
// ---------------------------------------------------------------------------

function applyPrescriptFixes(tempApisDir) {
  const files = walkSync(tempApisDir);
  let fixed = 0;

  for (const file of files) {
    if (!file.endsWith(".yaml") && !file.endsWith(".yml")) continue;

    let rawdata = fs.readFileSync(file, "utf8");
    let lines   = rawdata.split("\n");
    let out     = [];
    let changed = false;

    if (file.includes(path.join("transforms", "schemas", "transform.yaml")) ||
        file.includes(path.join("sources",    "schemas", "transform.yaml"))) {
      for (let line of lines) {
        if (line.includes("oneOf")) { line = line.replaceAll("oneOf:", "type: object"); changed = true; }
        if (!line.includes("- $ref:")) out.push(line);
      }
      lines = out; out = [];
    }

    if (file.includes(path.join("workflows", "schemas", "workflowtrigger.yaml"))) {
      for (let line of lines) {
        if (line.includes("anyOf")) { line = line.replaceAll("anyOf:", "type: object"); changed = true; }
        if (!line.includes("- $ref:")) out.push(line);
      }
      lines = out; out = [];
    }

    if (file.includes(path.join("search", "schemas", "searchdocument.yaml")) ||
        file.includes(path.join("search", "schemas", "searchdocuments.yaml"))) {
      lines = ["type: object"];
      changed = true;
    }

    if (changed) {
      fs.writeFileSync(file, lines.join("\n"), "utf8");
      fixed++;
    }
  }

  console.log(`  prescript: fixed ${fixed} file(s) in temp copy`);
}

// ---------------------------------------------------------------------------
// Bundle a single partition's openapi.yaml with redocly
// ---------------------------------------------------------------------------

function bundlePartition(partitionName, tempApisDir) {
  const inputSpec  = path.join(tempApisDir, partitionName, "openapi.yaml");
  const outputSpec = path.join(BUNDLED_DIR, `${partitionName}.yaml`);

  fs.mkdirSync(BUNDLED_DIR, { recursive: true });

  const result = spawnSync(
    "redocly",
    ["bundle", inputSpec, "-o", outputSpec, "--force"],
    { encoding: "utf8" }
  );

  return {
    ok:     result.status === 0,
    stdout: result.stdout || "",
    stderr: result.stderr || "",
    outputSpec,
  };
}

// ---------------------------------------------------------------------------
// Generate per-partition config YAML
// ---------------------------------------------------------------------------

function writePartitionConfig(partitionName) {
  const pascal = partitionName
    .split("-")
    .map(s => s.charAt(0).toUpperCase() + s.slice(1))
    .join("");

  const packageName   = `api_${partitionName.replaceAll("-", "_")}_v1`;
  const subModuleName = `${pascal}V1`;
  const importPath    = `github.com/sailpoint-oss/golang-sdk/v2/${packageName}`;

  const config = [
    `templateDir: ${TEMPLATE_DIR}`,
    `packageName: ${packageName}`,
    `packageVersion: 1.0.0`,
    `generateInterfaces: false`,
    `disallowAdditionalPropertiesIfNotPresent: false`,
    `useOneOfDiscriminatorLookup: false`,
    `enumClassPrefix: true`,
    `gitRepoId: ${GIT_REPO_ID}`,
    `gitUserId: ${GIT_USER_ID}`,
    `isGoSubmodule: true`,
    `apiVersion: ${API_VERSION}`,
    `funtionPrefix: ${FN_PREFIX}`,
    `subModuleName: ${subModuleName}`,
    `moduleImportPath: ${importPath}`,
    `files:`,
    `  developerSite_code_examples.mustache:`,
    `    templateType: APIDocs`,
    `    destinationFilename: developerSite_code_examples.yaml`,
    `  docs_methods_index.mustache:`,
    `    templateType: SupportingFiles`,
    `    destinationFilename: docs/Methods/Index.md`,
    `  docs_models_index.mustache:`,
    `    templateType: SupportingFiles`,
    `    destinationFilename: Index.md`,
  ].join("\n");

  const configPath = path.join(TEMP_DIR, `${partitionName}-config.yaml`);
  fs.writeFileSync(configPath, config, "utf8");
  return configPath;
}

// ---------------------------------------------------------------------------
// Run openapi-generator for a single partition (output captured)
// ---------------------------------------------------------------------------

function generatePartition(partitionName, bundledSpec, configPath) {
  const packageName = `api_${partitionName.replaceAll("-", "_")}_v1`;
  const outputDir   = path.join(SDK_ROOT, packageName);

  if (fs.existsSync(outputDir)) {
    fs.rmSync(outputDir, { recursive: true, force: true });
  }

  const result = spawnSync(
    "java",
    [
      "-jar", JAR,
      "generate",
      "-i", bundledSpec,
      "-g", "go",
      "-o", outputDir,
      "--global-property", "skipFormModel=false,apiDocs=true,modelDocs=true",
      "--config", configPath,
    ],
    { encoding: "utf8" }
  );

  return {
    ok:        result.status === 0,
    stdout:    result.stdout || "",
    stderr:    result.stderr || "",
    outputDir,
  };
}

// ---------------------------------------------------------------------------
// Run postscript.js on the generated output (output captured)
// ---------------------------------------------------------------------------

function runPostscript(outputDir) {
  const result = spawnSync(
    "node",
    [POSTSCRIPT, outputDir],
    { encoding: "utf8" }
  );

  return {
    ok:     result.status === 0,
    stdout: result.stdout || "",
    stderr: result.stderr || "",
  };
}

// ---------------------------------------------------------------------------
// Error logging
// ---------------------------------------------------------------------------

/**
 * Collect the key YAML spec files for a partition so the AI report is
 * self-contained. Files larger than MAX_FILE_BYTES are truncated with a note.
 */
const MAX_FILE_BYTES = 20_000;

function collectSpecFiles(partitionName, tempApisDir) {
  const partDir = path.join(tempApisDir, partitionName);
  if (!fs.existsSync(partDir)) return {};

  const collected = {};
  const files = walkSync(partDir).filter(f => f.endsWith(".yaml"));

  for (const f of files) {
    const relPath = path.relative(path.join(tempApisDir, ".."), f);
    let content = readFileSafe(f);
    if (Buffer.byteLength(content, "utf8") > MAX_FILE_BYTES) {
      content = content.slice(0, MAX_FILE_BYTES) + "\n\n[... truncated — file exceeds 20 KB ...]";
    }
    collected[relPath] = content;
  }

  return collected;
}

function writeErrorReport(partitionName, step, errorOutput, tempApisDir, apisSourceDir) {
  fs.mkdirSync(ERROR_DIR, { recursive: true });

  const specFiles   = collectSpecFiles(partitionName, path.join(tempApisDir, "apis"));
  const sourceDir   = path.join(apisSourceDir, partitionName);
  const reportPath  = path.join(ERROR_DIR, `${partitionName}-error.md`);

  const fileBlocks = Object.entries(specFiles)
    .map(([relPath, content]) => `### \`${relPath}\`\n\`\`\`yaml\n${content}\n\`\`\``)
    .join("\n\n");

  const report = `# SDK Build Error: \`${partitionName}\`

## Context for AI

This file is a self-contained error report for the \`${partitionName}\` API partition.
It contains the build error and all relevant OpenAPI spec files needed to diagnose and fix the problem.

**Source directory to fix:** \`${sourceDir}\`
**Failed step:** ${step}

---

## Error Output

\`\`\`
${errorOutput.trim()}
\`\`\`

---

## Fix Instructions

1. Read the error output above carefully.
2. Identify which spec file(s) below are causing the problem.
3. Apply fixes directly to the source files under \`${sourceDir}\`.
4. Do **not** edit files in \`.sdk-build-tmp/\` — those are generated copies.
5. After fixing, re-run the build for just this partition:
   \`\`\`bash
   node sdk-resources/build-versioned-sdk.js <path-to-apis> --partition ${partitionName}
   \`\`\`

---

## Spec Files

${fileBlocks || "_No spec files found._"}
`;

  fs.writeFileSync(reportPath, report, "utf8");
  return reportPath;
}

function writeSummaryReport(results, apisSourceDir) {
  fs.mkdirSync(ERROR_DIR, { recursive: true });

  const failureLines = results.failed.map(({ partition, step, reportPath }) =>
    `| \`${partition}\` | ${step} | [${path.basename(reportPath)}](./${path.basename(reportPath)}) |`
  ).join("\n");

  const summary = `# SDK Build Error Summary

**Date:** ${new Date().toISOString()}
**APIs directory:** \`${apisSourceDir}\`
**Total partitions:** ${results.total}
**Succeeded:** ${results.success.length}
**Failed:** ${results.failed.length}

## Failed Partitions

| Partition | Failed Step | Error Report |
|-----------|-------------|--------------|
${failureLines || "_(none)_"}

## How to Fix

Each error report in this directory is self-contained and can be given directly to an AI.
It includes the full error output and all relevant spec file contents.

Fix partitions one at a time:
\`\`\`bash
# Fix a single partition
node sdk-resources/build-versioned-sdk.js <path-to-apis> --partition <partition-name>

# Re-run all after fixes
node sdk-resources/build-versioned-sdk.js <path-to-apis>
\`\`\`
`;

  const summaryPath = path.join(ERROR_DIR, "summary.md");
  fs.writeFileSync(summaryPath, summary, "utf8");
  return summaryPath;
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

async function main() {
  if (!fs.existsSync(apisDir)) {
    console.error(`Error: apis directory not found: ${apisDir}`);
    process.exit(1);
  }

  if (!fs.existsSync(JAR)) {
    console.error(`Error: openapi-generator-cli.jar not found at ${JAR}`);
    console.error("  Download it with:");
    console.error("  wget -q https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.12.0/openapi-generator-cli-7.12.0.jar -O openapi-generator-cli.jar");
    process.exit(1);
  }

  const allPartitions = fs.readdirSync(apisDir, { withFileTypes: true })
    .filter(e => e.isDirectory() && e.name !== "shared")
    .map(e => e.name)
    .sort();

  const partitions = onlyPartition
    ? allPartitions.filter(p => p === onlyPartition)
    : allPartitions;

  if (partitions.length === 0) {
    console.error(`No partitions found${onlyPartition ? ` matching '${onlyPartition}'` : ""} in ${apisDir}`);
    process.exit(1);
  }

  console.log(`\nFound ${partitions.length} partition(s) to build\n`);

  // Set up temp directory
  console.log(`[SETUP] Copying apis/ → ${path.relative(SDK_ROOT, TEMP_DIR)}/`);
  if (fs.existsSync(TEMP_DIR)) fs.rmSync(TEMP_DIR, { recursive: true, force: true });
  copyDirSync(apisDir, path.join(TEMP_DIR, "apis"));

  console.log("[SETUP] Applying prescript fixes to temp copy ...");
  applyPrescriptFixes(path.join(TEMP_DIR, "apis"));

  // Clear any previous error reports
  if (fs.existsSync(ERROR_DIR)) fs.rmSync(ERROR_DIR, { recursive: true, force: true });

  const results = {
    total:   partitions.length,
    success: [],
    failed:  [],
  };

  for (const partition of partitions) {
    console.log(`\n${"=".repeat(60)}`);
    console.log(`  Building: ${partition}`);
    console.log(`${"=".repeat(60)}`);

    // --- Step 1: Bundle ---
    console.log("  [1/4] Bundling spec ...");
    const bundle = bundlePartition(partition, path.join(TEMP_DIR, "apis"));
    if (!bundle.ok) {
      const errorOutput = [bundle.stdout, bundle.stderr].filter(Boolean).join("\n");
      console.error(`  ✗ bundling failed`);
      const reportPath = writeErrorReport(partition, "bundling", errorOutput, TEMP_DIR, apisDir);
      results.failed.push({ partition, step: "bundling", reportPath });
      continue;
    }

    // --- Step 2: Config ---
    console.log("  [2/4] Writing generator config ...");
    const configPath = writePartitionConfig(partition);

    // --- Step 3: Generate ---
    console.log("  [3/4] Generating Go SDK ...");
    const gen = generatePartition(partition, bundle.outputSpec, configPath);
    if (!gen.ok) {
      const errorOutput = [gen.stdout, gen.stderr].filter(Boolean).join("\n");
      console.error(`  ✗ generation failed`);
      const reportPath = writeErrorReport(partition, "generation", errorOutput, TEMP_DIR, apisDir);
      results.failed.push({ partition, step: "generation", reportPath });
      continue;
    }

    // --- Step 4: Postscript ---
    console.log("  [4/4] Running postscript ...");
    const post = runPostscript(gen.outputDir);
    if (!post.ok) {
      const errorOutput = [post.stdout, post.stderr].filter(Boolean).join("\n");
      console.error(`  ✗ postscript failed`);
      const reportPath = writeErrorReport(partition, "postscript", errorOutput, TEMP_DIR, apisDir);
      results.failed.push({ partition, step: "postscript", reportPath });
      continue;
    }

    results.success.push(partition);
    console.log(`  ✓ ${partition} → ${path.relative(SDK_ROOT, gen.outputDir)}/`);
  }

  // Cleanup
  if (!keepTmp) {
    console.log("\n[CLEANUP] Removing .sdk-build-tmp/ ...");
    fs.rmSync(TEMP_DIR, { recursive: true, force: true });
  }

  // Regenerate client.go to include all current api_*_v1 packages
  console.log("\n[CLIENT] Regenerating client.go ...");
  generateClientGo();

  // Wire local submodules via go.work (avoids fetching unpublished packages from GitHub)
  console.log("\n[MODULES] Syncing Go workspace ...");
  const syncResult = spawnSync("node", [path.join(__dirname, "sync-go-modules.js")], {
    cwd: SDK_ROOT,
    encoding: "utf8",
    stdio: "inherit",
  });
  if (syncResult.status !== 0) {
    console.error("Go module sync failed. Run: node sdk-resources/sync-go-modules.js");
    process.exit(1);
  }

  // Write error reports
  if (results.failed.length > 0) {
    const summaryPath = writeSummaryReport(results, apisDir);
    console.log(`\n[ERRORS] ${results.failed.length} partition(s) failed.`);
    console.log(`  Summary:  ${path.relative(SDK_ROOT, summaryPath)}`);
    console.log(`  Reports:  ${path.relative(SDK_ROOT, ERROR_DIR)}/`);
    console.log(`\n  Failed partitions:`);
    for (const { partition, step, reportPath } of results.failed) {
      console.log(`    ✗ ${partition} (${step}) → ${path.relative(SDK_ROOT, reportPath)}`);
    }
  }

  // Summary
  console.log("\n" + "=".repeat(60));
  console.log("BUILD SUMMARY");
  console.log("=".repeat(60));
  console.log(`  Success: ${results.success.length} / ${results.total}`);
  console.log(`  Failed:  ${results.failed.length} / ${results.total}`);

  if (results.failed.length > 0) {
    process.exit(1);
  }
}

/**
 * Convert a snake_case package name like api_access_profiles_v1 to a
 * PascalCase Go field name like AccessProfiles.
 */
function packageToFieldName(pkgName) {
  return pkgName
    .replace(/^api_/, "")
    .replace(/_v\d+$/, "")
    .split("_")
    .map(w => w.charAt(0).toUpperCase() + w.slice(1))
    .join("");
}

/**
 * Regenerate golang-sdk/client.go from the api_*_v1 packages currently on disk.
 * Also keeps api_generic, api_nerm, and api_nerm_v2025 as special cases.
 */
function generateClientGo() {
  const MODULE = "github.com/sailpoint-oss/golang-sdk/v2";

  // Discover all api_*_v1 packages on disk
  const partitionPkgs = fs.readdirSync(SDK_ROOT)
    .filter(d => /^api_.+_v\d+$/.test(d) && fs.statSync(path.join(SDK_ROOT, d)).isDirectory())
    .sort();

  if (partitionPkgs.length === 0) {
    console.log("  No api_*_v1 packages found, skipping client.go generation.");
    return;
  }

  // Build import block
  const partitionImports = partitionPkgs
    .map(pkg => `\t${pkg} "${MODULE}/${pkg}"`)
    .join("\n");

  // Build struct fields
  const partitionFields = partitionPkgs
    .map(pkg => `\t${packageToFieldName(pkg)} *${pkg}.APIClient`)
    .join("\n");

  // Build NewAPIClient instantiation lines
  const partitionInits = partitionPkgs.map(pkg => {
    const field = packageToFieldName(pkg);
    return [
      `\t_cfg${field} := ${pkg}.NewConfiguration(`,
      `\t\tcfg.ClientConfiguration.ClientId,`,
      `\t\tcfg.ClientConfiguration.ClientSecret,`,
      `\t\tcfg.ClientConfiguration.BaseURL,`,
      `\t\tcfg.ClientConfiguration.TokenURL,`,
      `\t\tcfg.ClientConfiguration.Token,`,
      `\t\tconsumerSuffix,`,
      `\t\tcfg.Experimental,`,
      `\t)`,
      `\t_cfg${field}.HTTPClient = cfg.HTTPClient`,
      `\tc.${field} = ${pkg}.NewAPIClient(_cfg${field})`,
    ].join("\n");
  }).join("\n\n");

  const content = `/*
SailPoint Identity Security Cloud API

Use these APIs to interact with the SailPoint Identity Security Cloud platform.
We encourage you to join the SailPoint Developer Community forum at
https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

Per-resource versioned API packages (e.g. api_entitlements_v1, api_accounts_v1) are
generated automatically — this file is regenerated by build-versioned-sdk.js after
each build, so new partitions appear here without any manual changes.
*/

// Code generated by build-versioned-sdk.js; DO NOT EDIT.

package sailpoint

import (
\t"regexp"

\t"github.com/hashicorp/go-retryablehttp"
${partitionImports}
\tgeneric   "${MODULE}/api_generic"
\tnerm      "${MODULE}/api_nerm"
\tnermv2025 "${MODULE}/api_nerm_v2025"
)

var (
\tjsonCheck = regexp.MustCompile(\`(?i:(?:application|text)/(?:vnd\\.[^;]+\\+)?json)\`)
\txmlCheck  = regexp.MustCompile(\`(?i:(?:application|text)/xml)\`)
)

// APIClient manages communication with the SailPoint Identity Security Cloud API.
// In most cases there should be only one, shared, APIClient.
//
// Per-resource versioned clients are available as fields (e.g. c.Accounts, c.Roles).
// Use NewPartitionConfiguration to build a config for directly importing a single package.
type APIClient struct {
\tcfg *Configuration

\t// Per-resource versioned clients (auto-generated)
${partitionFields}

\t// Shared utility clients
\tGeneric   *generic.APIClient
\tNERM      *nerm.APIClient
\tNERMV2025 *nermv2025.APIClient

\ttoken string
}

// NewAPIClient creates a new API client from the provided configuration.
func NewAPIClient(cfg *Configuration) *APIClient {
\tif cfg.HTTPClient == nil {
\t\tcfg.HTTPClient = retryablehttp.NewClient()
\t}

\tc := &APIClient{}

\tconsumerSuffix := cfg.ConsumerSuffix()

\t// Per-resource versioned clients
${partitionInits}

\t// Shared utility clients
\tCVGeneric := generic.NewConfiguration(
\t\tcfg.ClientConfiguration.ClientId,
\t\tcfg.ClientConfiguration.ClientSecret,
\t\tcfg.ClientConfiguration.BaseURL,
\t\tcfg.ClientConfiguration.TokenURL,
\t\tcfg.ClientConfiguration.Token,
\t\tconsumerSuffix,
\t\tcfg.Experimental,
\t)
\tCNERM := nerm.NewConfiguration(
\t\tcfg.ClientConfiguration.ClientId,
\t\tcfg.ClientConfiguration.ClientSecret,
\t\tcfg.ClientConfiguration.NermBaseURL,
\t\tcfg.ClientConfiguration.TokenURL,
\t\tcfg.ClientConfiguration.Token,
\t\tconsumerSuffix,
\t\tcfg.Experimental,
\t)
\tCNERMV2025 := nermv2025.NewConfiguration(
\t\tcfg.ClientConfiguration.ClientId,
\t\tcfg.ClientConfiguration.ClientSecret,
\t\tcfg.ClientConfiguration.NermBaseURL+"/v2025",
\t\tcfg.ClientConfiguration.TokenURL,
\t\tcfg.ClientConfiguration.Token,
\t\tconsumerSuffix,
\t\tcfg.Experimental,
\t)

\tCVGeneric.HTTPClient = cfg.HTTPClient
\tCNERM.HTTPClient = cfg.HTTPClient
\tCNERMV2025.HTTPClient = cfg.HTTPClient

\tc.Generic = generic.NewAPIClient(CVGeneric)
\tc.NERM = nerm.NewAPIClient(CNERM)
\tc.NERMV2025 = nermv2025.NewAPIClient(CNERMV2025)

\treturn c
}

// NewPartitionConfiguration returns the configuration parameters needed to
// instantiate any per-resource versioned API package directly.
// Example:
//
//\tcfg := api_accounts_v1.NewConfiguration(NewPartitionConfiguration(sailpointCfg))
//\tclient := api_accounts_v1.NewAPIClient(cfg)
func NewPartitionConfiguration(cfg *Configuration) (clientId, clientSecret, baseURL, tokenURL, token, consumerSuffix string, experimental bool) {
\treturn cfg.ClientConfiguration.ClientId,
\t\tcfg.ClientConfiguration.ClientSecret,
\t\tcfg.ClientConfiguration.BaseURL,
\t\tcfg.ClientConfiguration.TokenURL,
\t\tcfg.ClientConfiguration.Token,
\t\tcfg.ConsumerSuffix(),
\t\tcfg.Experimental
}
`
    .replace("${partitionImports}", partitionImports)
    .replace("${partitionFields}", partitionFields)
    .replace("${partitionInits}", partitionInits)
    .replaceAll("${MODULE}", MODULE);

  const outPath = path.join(SDK_ROOT, "client.go");
  fs.writeFileSync(outPath, content, "utf8");
  console.log(`  Wrote client.go with ${partitionPkgs.length} partition client(s) + Generic/NERM/NERMV2025`);
}

main().catch(err => {
  console.error(`Unexpected error: ${err.message}`);
  process.exit(1);
});
