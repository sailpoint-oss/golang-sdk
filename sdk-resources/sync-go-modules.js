#!/usr/bin/env node
/**
 * sync-go-modules.js
 *
 * Prepares the multi-module golang-sdk workspace for local builds:
 *   1. Ensures each api_* package has a go.mod with the correct module path
 *   2. Creates/updates go.work to wire the root module and all submodules together
 *   3. Runs go work sync and go mod tidy on the root module
 *
 * Without go.work, `go build` resolves submodules from the GitHub proxy — which fails
 * for packages that exist only in this repo.
 *
 * Usage:
 *   node sdk-resources/sync-go-modules.js
 */

"use strict";

const fs = require("fs");
const path = require("path");
const { spawnSync } = require("child_process");

const SDK_ROOT = path.resolve(__dirname, "..");
const MODULE_PREFIX = "github.com/sailpoint-oss/golang-sdk/v3";

const SUBMODULE_GO_TEMPLATE = (modulePath) => `module ${modulePath}

go 1.18

require (
	github.com/hashicorp/go-retryablehttp v0.7.2
)
`;

function discoverSubmoduleDirs() {
  return fs.readdirSync(SDK_ROOT)
    .filter(name => {
      if (!/^api_/.test(name)) return false;
      const dir = path.join(SDK_ROOT, name);
      return fs.statSync(dir).isDirectory();
    })
    .sort();
}

function ensureSubmoduleGoMod(dirName) {
  const modulePath = `${MODULE_PREFIX}/${dirName}`;
  const goModPath = path.join(SDK_ROOT, dirName, "go.mod");
  let needsWrite = true;

  if (fs.existsSync(goModPath)) {
    const content = fs.readFileSync(goModPath, "utf8");
    if (content.startsWith(`module ${modulePath}`)) {
      needsWrite = false;
    }
  }

  if (needsWrite) {
    fs.writeFileSync(goModPath, SUBMODULE_GO_TEMPLATE(modulePath), "utf8");
    console.log(`  Wrote ${dirName}/go.mod → ${modulePath}`);
  }
}

function run(cmd, args, label) {
  const result = spawnSync(cmd, args, { cwd: SDK_ROOT, encoding: "utf8" });
  if (result.status !== 0) {
    const out = [result.stdout, result.stderr].filter(Boolean).join("\n");
    console.error(`  ✗ ${label} failed:\n${out}`);
    process.exit(1);
  }
  return result;
}

function syncGoWorkspace() {
  const submodules = discoverSubmoduleDirs();
  if (submodules.length === 0) {
    console.log("No api_* submodule directories found.");
    return;
  }

  console.log(`[MODULES] Fixing go.mod for ${submodules.length} submodule(s) ...`);
  for (const dir of submodules) {
    ensureSubmoduleGoMod(dir);
  }

  const workPath = path.join(SDK_ROOT, "go.work");
  if (!fs.existsSync(workPath)) {
    console.log("[WORKSPACE] Initializing go.work ...");
    run("go", ["work", "init", "."], "go work init");
  }

  console.log("[WORKSPACE] Registering all modules (go work use -r .) ...");
  run("go", ["work", "use", "-r", "."], "go work use");

  console.log("[WORKSPACE] Syncing workspace (go work sync) ...");
  run("go", ["work", "sync"], "go work sync");

  // Skip `go mod tidy` on the root module: unpublished submodules are not on the
  // module proxy yet, so tidy would try to fetch them from GitHub and fail.
  // go.work wires local modules; use `go build .` from the repo root.

  console.log("Go workspace ready. Build with: go build .");
}

syncGoWorkspace();
