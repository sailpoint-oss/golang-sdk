# Versioned SDK Migration Process

This document captures every task performed in the Go SDK to migrate from the monolithic
year/stage-based versioning (`api_v3`, `api_beta`, `api_v2024`, `api_v2025`, `api_v2026`) to
per-resource versioned packages (`api_{resource}_v1`). Use it as a template for applying the
same changes to the Python, TypeScript, and PowerShell SDKs.

---

## Background: What Changed and Why

**Before:** One large SDK package was generated per API "version" (v3, beta, v2024, v2025,
v2026). Every resource (accounts, roles, workflows, …) lived inside the same package. Adding a
new API endpoint required regenerating the entire version package and shipping a monolithic update.

**After:** One small SDK package is generated per API **resource partition** (accounts, roles,
workflows, …), all at a single opinionated version (`v1`). The build script auto-discovers
partitions from the `apis/` directory structure in `api-specs`, so new endpoints appear
automatically without any script or config changes.

---

## Task List

### 1. Create `sdk-resources/build-versioned-sdk.js`

This is the core of the migration. The script replaces all of the old per-version build steps.

**What it does:**

1. Accepts a path to the `apis/` directory in the `api-specs` repository as its first argument.
2. Reads all subdirectories under `apis/` (except `shared/`) as **partitions** — one partition
   per resource group (e.g., `accounts`, `roles`, `workflows`).
3. Copies `apis/` to a git-ignored temp dir (`.sdk-build-tmp/`) so source files stay clean.
4. Applies **prescript fixes** to the temp copy (YAML workarounds for known generator issues —
   previously a separate `prescript.js` run against the source directory).
5. Bundles each partition's `openapi.yaml` with **Redocly** to resolve `$ref` pointers.
6. Dynamically generates an OpenAPI Generator config YAML for each partition (package name,
   module import path, template directory, etc.) — eliminating one static config file per version.
7. Runs `openapi-generator-cli` to generate the SDK package into `api_{resource}_v1/`.
8. Runs `postscript.js` on each generated package.
9. On any failure, writes a self-contained error report to `build-errors/{partition}-error.md`
   (includes the spec YAML and fix instructions, suitable to hand to an AI).
10. After all partitions are built, **regenerates `client.go`** from all `api_*_v1` packages
    found on disk — new partitions appear in the top-level client automatically.
11. Calls `sync-go-modules.js` (Go-specific) to wire the local workspace.

**CLI usage:**

```bash
# Build all partitions
node sdk-resources/build-versioned-sdk.js <path-to-apis-dir>

# Build a single partition (fast iteration)
node sdk-resources/build-versioned-sdk.js <path-to-apis-dir> --partition accounts

# Keep the temp build directory for inspection
node sdk-resources/build-versioned-sdk.js <path-to-apis-dir> --keep-tmp
```

**Language-specific notes for other SDKs:**

- The script is written in Node.js and uses only the built-in `fs`, `path`, and `child_process`
  modules — no `npm install` required. Adapt the logic to the tooling available in each SDK repo
  if a Node-based script doesn't fit.
- The config YAML generation block (`writePartitionConfig`) is the most language-specific part.
  Replace the Go-specific options (`isGoSubmodule`, `gitRepoId`, etc.) with the appropriate
  openapi-generator options for the target language.
- The `client.go` auto-generation step (`generateClientGo`) is entirely Go-specific. For Python,
  TypeScript, and PowerShell, add an equivalent step that builds a top-level re-export module or
  barrel file from the discovered partition packages.
- The prescript logic (inline in this script) was previously a standalone `prescript.js` that
  mutated source files directly. Moving it to operate on the temp copy means source files are
  never modified during a build.

---

### 2. Create `sdk-resources/sync-go-modules.js` _(Go-specific)_

Manages the Go multi-module workspace (`go.work`) required because each `api_*_v1` package is
its own Go module.

**What it does:**

1. Discovers all `api_*` directories under the SDK root.
2. Ensures each has a `go.mod` with the correct module path.
3. Creates or updates `go.work` to include all submodules.
4. Runs `go work sync`.

**Language-specific notes:**

- Python: equivalent is ensuring each partition package has the correct `setup.py` /
  `pyproject.toml` and that a `requirements.txt` or workspace file wires them together for local
  development.
- TypeScript: equivalent is ensuring each partition package has a `package.json` and that the
  root `package.json` uses `workspaces` or a `tsconfig.json` `paths` mapping.
- PowerShell: no direct equivalent needed; the generator typically outputs a flat module
  structure. Validate that the generated manifest (`.psd1`) references all exported functions.

---

### 3. Remove Old Version-Specific SDK Packages

Delete the following directories from the repository:

| Directory    | Reason                                     |
|--------------|--------------------------------------------|
| `api_v3/`    | Replaced by per-resource `api_*_v1/` packages |
| `api_beta/`  | Replaced by per-resource `api_*_v1/` packages |
| `api_v2024/` | Replaced by per-resource `api_*_v1/` packages |
| `api_v2025/` | Replaced by per-resource `api_*_v1/` packages |
| `api_v2026/` | Replaced by per-resource `api_*_v1/` packages |

Keep `api_generic/`, `api_nerm/`, and `api_nerm_v2025/` — these are not partition-based and
continue to be built by separate steps.

**Language-specific notes:** Apply to whichever equivalent directories exist in the other SDK
repos (e.g., `sailpoint/beta/`, `sailpoint/v3/`, etc. in Python).

---

### 4. Remove Old Version-Specific Config YAMLs from `sdk-resources/`

Delete the static per-version config files that `build-versioned-sdk.js` now generates
dynamically:

- `sdk-resources/v3-config.yaml`
- `sdk-resources/beta-config.yaml`
- `sdk-resources/v2024-config.yaml`
- `sdk-resources/v2025-config.yaml`
- `sdk-resources/v2026-config.yaml`

Keep `sdk-resources/generic-config.yaml`, `sdk-resources/nerm-config.yaml`, and
`sdk-resources/v2025-nerm-config.yaml` — NERM and Generic are still built with static configs.

---

### 5. Remove `sdk-resources/prescript.js` as a Standalone Script

The prescript YAML-fix logic has been moved **inline** into `build-versioned-sdk.js`
(`applyPrescriptFixes` function). It now runs against the temp copy of `apis/` rather than
against the source files directly.

If your other SDK repos have a standalone `prescript.js` (or equivalent), migrate its logic into
the new build script in the same way.

---

### 6. Update `sdk-resources/versioned-config.yaml`

Ensure the shared `versioned-config.yaml` (template settings, additional properties) reflects
any options that were previously spread across the individual version config files. The
`build-versioned-sdk.js` `writePartitionConfig` function generates per-partition configs that
should include all required fields.

---

### 7. Update `.github/actions/build-sdk/action.yaml`

**Replace** all individual per-version build steps with a single versioned build step:

```yaml
# OLD (remove all of these)
- name: Run Prescript
  run: node sdk-resources/prescript.js ${{ inputs.api-specs-path }}/idn

- name: Build V3 SDK
  run: |
    rm -rf ./api_v3
    java -jar openapi-generator-cli.jar generate \
      -i ${{ inputs.api-specs-path }}/idn/sailpoint-api.v3.yaml \
      ...

- name: Build Beta SDK
  ...

- name: Build V2024 SDK
  ...

- name: Build V2025 SDK
  ...

- name: Build V2026 SDK
  ...

# NEW (replace with this single step)
- name: Build Versioned SDK (per-partition)
  run: |
    node sdk-resources/build-versioned-sdk.js \
      ${{ inputs.api-specs-path }}/idn/src/main/yaml/apis
```

Note the input path change: from `${{ inputs.api-specs-path }}/idn/sailpoint-api.v*.yaml`
(monolithic spec files) to `${{ inputs.api-specs-path }}/idn/src/main/yaml/apis` (the `apis/`
directory containing one subdirectory per partition).

**Keep** the Generic SDK, NERM SDK, and NERM V2025 SDK build steps — they are unchanged.

**Add** a "Wire Go workspace" step after all SDK builds and before the test step _(Go-specific)_:

```yaml
- name: Wire Go workspace
  run: |
    go work init . 2>/dev/null || true
    go work use -r .
    go work sync
```

**Update** the test step to use the workspace-aware build command _(Go-specific)_:

```yaml
# OLD
- name: Build and Test
  run: |
    go get -d ./...
    go install
    go test

# NEW
- name: Build and Test
  run: |
    go build -o /dev/null .
    if [ "${{ inputs.skip-tests }}" != "true" ]; then
      go test ./...
    fi
```

**Update the Node.js version** (minor but noted): the action previously specified Node 24; set
it to Node 20 (LTS) for broader compatibility unless the repo has a specific requirement.

**Language-specific notes for other SDKs:** Substitute the Go workspace wiring and `go build`/
`go test` commands with the equivalent for the target language (e.g., `pip install -e .` +
`pytest` for Python, `npm install` + `npm test` for TypeScript, `Import-Module` + `Pester` for
PowerShell).

---

### 8. Update `.github/workflows/push_sdk_docs_to_dev_portal.yaml`

**Replace** the hardcoded per-version `rsync` calls with a dynamic loop over `api_*_v1`
directories:

```yaml
# OLD (hardcoded per version)
- name: Sync files between folders
  run: |
    rsync -cav --delete golang-sdk/api_v3/docs/Methods developer-community/docs/tools/sdk/go/Reference/V3
    rsync -cav --delete golang-sdk/api_v3/docs/Models  developer-community/docs/tools/sdk/go/Reference/V3
    rsync -av  golang-sdk/api_v3/docs/Examples/go_code_examples_overlay.yaml developer-community/static/code-examples/v3/
    # ... repeated for beta, v2024, v2025, v2026 ...

# NEW (dynamic loop)
- name: Sync files between folders
  run: |
    for partition_dir in golang-sdk/api_*_v1; do
      [ -d "$partition_dir" ] || continue
      resource=$(basename "$partition_dir")
      rsync -cav --delete "$partition_dir/docs/Methods" "developer-community/docs/tools/sdk/go/Reference/$resource"
      rsync -cav --delete "$partition_dir/docs/Models"  "developer-community/docs/tools/sdk/go/Reference/$resource"
      if [ -f "$partition_dir/docs/Examples/go_code_examples_overlay.yaml" ]; then
        rsync -av "$partition_dir/docs/Examples/go_code_examples_overlay.yaml" \
          "developer-community/static/code-examples/$resource/"
      fi
    done
```

This ensures new partitions are synced automatically without any workflow changes.

---

### 9. Update `Makefile` (or Equivalent Build Script)

**Replace** the old per-version build invocations with the new versioned build script:

```makefile
# OLD
build:
    node sdk-resources/prescript.js api-specs/idn
    rm -rf ./api_v3
    java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v3.yaml ...
    # ... repeated for beta, v2024, v2025, v2026 ...

# NEW
APIS_DIR ?= api-specs/idn/src/main/yaml/apis

build:
    node sdk-resources/build-versioned-sdk.js $(APIS_DIR)
    # Generic and NERM builds remain here, unchanged
```

**Add** a single-partition build target for fast iteration during development:

```makefile
build-partition:
    node sdk-resources/build-versioned-sdk.js $(APIS_DIR) --partition $(PARTITION)
```

**Add** a module-sync target _(Go-specific)_:

```makefile
sync-modules:
    node sdk-resources/sync-go-modules.js
```

**Update** the `test` target to use workspace-aware commands _(Go-specific)_:

```makefile
# OLD
test:
    go get -d ./...
    go install
    go test

# NEW
test: sync-modules
    go build -o /dev/null .
    go test ./...
```

---

### 10. Update `client.go` (Top-Level SDK Entry Point) _(Go-specific)_

The root `client.go` used to be hand-maintained with one field per API version
(`V3 *v3.APIClient`, `Beta *beta.APIClient`, etc.).

It is now **auto-generated** by `build-versioned-sdk.js` (`generateClientGo` function) after
every build. The generated file:

- Imports all `api_*_v1` packages found on disk.
- Exposes each as a typed field on `APIClient` (e.g., `c.Accounts`, `c.Workflows`).
- Adds the header comment `// Code generated by build-versioned-sdk.js; DO NOT EDIT.`
- Adds the helper `NewPartitionConfiguration` for users who want to instantiate a single
  partition package directly.

**Do not hand-edit this file after migration** — it will be overwritten on the next build.

**Language-specific notes:** For Python, this is equivalent to updating the top-level
`__init__.py` to import from all partition packages. For TypeScript, update `index.ts` to
re-export all partition modules. For PowerShell, update the root `.psm1` or barrel module.
Consider whether to auto-generate these or maintain them manually — auto-generation is strongly
preferred.

---

### 11. Update Tests

> **⚠️ CRITICAL — do NOT delete tests.**
> Every existing test must be kept and migrated to its V1 equivalent. Deleting tests removes
> coverage that is hard to recover. If the new API shape makes a test temporarily impossible to
> pass, comment it out with a `// TODO partition-strategy:` note explaining why, but never
> simply remove it.

**Replace** imports of monolithic version packages with imports of specific per-resource
partition packages:

```go
// OLD
import (
    beta  "github.com/sailpoint-oss/golang-sdk/v2/api_beta"
    v2024 "github.com/sailpoint-oss/golang-sdk/v2/api_v2024"
    v3    "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
)

// NEW
import (
    "github.com/sailpoint-oss/golang-sdk/v2/api_accounts_v1"
    "github.com/sailpoint-oss/golang-sdk/v2/api_transforms_v1"
    "github.com/sailpoint-oss/golang-sdk/v2/api_workflows_v1"
)
```

**Update API call patterns** to use the new per-partition client fields and the `V1`-suffixed
method names generated by the new build:

```go
// OLD
apiClient.V3.AccountsAPI.ListAccounts(ctx).Execute()

// NEW
apiClient.Accounts.AccountsAPI.ListAccountsV1(ctx).Execute()
```

**Add tests for the direct partition client pattern** (using `NewPartitionConfiguration`):

```go
cfg := api_accounts_v1.NewConfiguration(NewPartitionConfiguration(configuration))
client := api_accounts_v1.NewAPIClient(cfg)
client.AccountsAPI.ListAccountsV1(ctx).Execute()
```

**Migration checklist for each existing test:**

1. Identify the old version package it imported (`v3`, `beta`, `v2024`, `v2025`, `v2026`).
2. Find the equivalent partition package (e.g. `api_accounts_v1` for anything that called an
   Accounts endpoint).
3. Append `V1` to every method name (`ListAccounts` → `ListAccountsV1`).
4. Update the client field path on `apiClient` (`apiClient.V3.AccountsAPI` →
   `apiClient.Accounts.AccountsAPI`).
5. Keep the assertion logic unchanged — only the import path and method name change.

---

### 12. Update `go.mod` _(Go-specific)_

The root `go.mod` no longer needs to reference per-version sub-packages as dependencies because
they are now separate Go modules wired via `go.work`. Remove any such entries and run
`go mod tidy` to clean up indirect dependencies that were only needed by the old monolithic
packages.

---

### 13. Add `go.work` and `go.work.sum` _(Go-specific)_

A `go.work` file is required to wire all `api_*_v1` submodules together for local development
and CI builds. It is generated by `sync-go-modules.js` and should be committed to the repository.

The workspace file lists every submodule under `use (...)`. It is regenerated automatically
after each build so it stays in sync as partitions are added or removed.

---

### 14. Update `.gitignore`

Add the following entries:

```gitignore
.sdk-build-tmp/
build-errors/
```

- `.sdk-build-tmp/` is the temp directory used during the build. It must not be committed.
- `build-errors/` is the per-partition error report directory. It is ephemeral build output and
  should not be committed.

---

### 15. Python — Generate Shared Root Support Files _(Python-specific)_

Each partition generated by the OpenAPI generator contains its own copies of four support
files: `api_client.py`, `api_response.py`, `rest.py`, and `exceptions.py`.  These files
are **structurally identical** across partitions — except that `api_client.py` hardcodes two
partition-specific references:

1. `import sailpoint.{partition}.models` — imported at module level.
2. `klass = getattr(sailpoint.{partition}.models, klass)` — used inside `__deserialize` to
   look up a response model class by string name.

This means you **cannot** simply re-export `ApiClient` from one partition and expect it to
work with another partition's API classes: when `SourcesApi` passes `'Source'` as the
response type, the client tries `getattr(sailpoint.accounts_v1.models, 'Source')`, which
raises `AttributeError` because `Source` only lives in `sailpoint.sources_v1.models`.

**Solution — copy and patch, not re-export:**

`generateInitPy()` in `build-versioned-sdk.js` should copy all four files from the first
discovered partition into `sailpoint/` and patch the partition-specific references on the fly:

| File | Patch required |
|------|---------------|
| `exceptions.py` | None — copy verbatim (no sailpoint imports). |
| `api_response.py` | None — copy verbatim (no sailpoint imports). |
| `rest.py` | `from sailpoint.{partition}.exceptions import` → `from sailpoint.exceptions import` |
| `api_client.py` | Fix all four partition-specific references (see below). |

**`api_client.py` patch steps:**

```text
1. from sailpoint.{partition}.api_response import  →  from sailpoint.api_response import
2. from sailpoint.{partition} import rest          →  from sailpoint import rest
3. from sailpoint.{partition}.exceptions import    →  from sailpoint.exceptions import
4. import sailpoint.{partition}.models             →  (remove this line entirely)
5. klass = getattr(sailpoint.{partition}.models, klass)
   →  replace with a sys.modules dynamic lookup (see below)
```

**Dynamic model lookup (step 5 replacement):**

```python
# Dynamic lookup across all loaded sailpoint.*.models
for _mod_name, _mod in list(sys.modules.items()):
    if (_mod_name.startswith("sailpoint.") and _mod_name.endswith(".models")
            and _mod is not None and hasattr(_mod, klass)):
        klass = getattr(_mod, klass)
        break
else:
    raise AttributeError(
        f"Cannot find model class {klass!r} in any loaded sailpoint partition module"
    )
```

This works because importing `SourcesApi` (for example) loads `sailpoint.sources_v1.models`
into `sys.modules` before any API call is made, so the right models module is always present
by the time deserialization runs.

Also ensure `import sys` is present at the top of the patched `api_client.py` (the generator
may not include it).

**Result:** users get a single, partition-agnostic `ApiClient`:

```python
from sailpoint import ApiClient, Configuration
client = ApiClient(Configuration())   # works with AccountsApi, SourcesApi, TransformsApi, …
```

---

### 16. Add a Grouped-Access Entry Point (Namespace / Module Object)

With per-partition packages, users would otherwise need a separate import for every resource they
use.  Each SDK should expose **one ergonomic entry point** that bundles all partition API classes
under a single name so a single import is enough for any combination of resources.

The pattern differs by language but the intent is the same: one import, all APIs available.

---

#### TypeScript — `SailPoint` namespace in `sdk-output/index.ts`

**Problem:** `export * from "./{partition}_v1/index"` across 100+ partitions causes TS2308
("already exported") because shared error models are inlined by Redocly into every partition's
`api.ts`. A flat `export *` therefore collides on model names.

**Solution:** Import each partition's API class with a private `_` alias, re-export it as a named
export for backward compatibility, and also expose it as a `const` inside a `SailPoint` namespace.
Using `const` (not `export { X }`) inside the namespace avoids TS1194 and TS2303.

```typescript
// --- Partition imports (private aliases avoid TS1194 / TS2303 in the namespace) ---
import { AccountsV1Api as _AccountsV1Api } from './accounts_v1/api';
import { TransformsV1Api as _TransformsV1Api } from './transforms_v1/api';
// ... one line per partition, generated by build-versioned-sdk.js

// --- Named exports (backward-compatible flat import style) ---
export { _AccountsV1Api as AccountsV1Api };
export { _TransformsV1Api as TransformsV1Api };
// ...

// --- SailPoint namespace (grouped access) ---
export namespace SailPoint {
  export const AccountsV1Api = _AccountsV1Api;
  export const TransformsV1Api = _TransformsV1Api;
  // ...
}
```

**Namespace naming convention — strip the version from the class name:**

The namespace exposes resources by name (`AccountsApi`), not by version (`AccountsV1Api`).
Method names carry the version suffix (`listAccountsV1`, `listAccountsV2`) so you always know
which version you are calling.  The import itself never changes when new versions land.

| Generated class | `SailPoint` namespace name |
|-----------------|---------------------------|
| `AccountsV1Api` | `SailPoint.AccountsApi`   |
| `AccountsV2Api` | `SailPoint.AccountsApi` (combined with v1) |
| `TransformsV1Api` | `SailPoint.TransformsApi` |

**Usage:**

```typescript
// Option A — named imports (existing / backward-compatible, version-explicit)
import { AccountsV1Api, Configuration } from 'sailpoint-api-client';
const api = new AccountsV1Api(new Configuration());
api.listAccountsV1({ limit: 10 });

// Option B — namespace (single import, resource-named, version-agnostic)
import { SailPoint, Configuration } from 'sailpoint-api-client';
const api = new SailPoint.AccountsApi(new Configuration());
api.listAccountsV1({ limit: 10 });  // today
api.listAccountsV2({ limit: 10 });  // when v2 partition lands — same import, no change

// Models are imported directly from the partition sub-path (not the root index)
import type { AccountV1 } from 'sailpoint-api-client/accounts_v1/api';
```

**Multi-version resources** (when both `accounts_v1` and `accounts_v2` exist):
`generateIndexTs()` automatically generates a combined class that extends the latest version
and copies older-version prototype methods via TypeScript interface merging + runtime prototype
copying.  Users see all versioned methods on the same `SailPoint.AccountsApi` instance with
full type safety.

`generateIndexTs()` in `build-versioned-sdk.js` produces `index.ts` automatically after every
build.  Do not hand-edit `index.ts`.

**TypeScript config:** Add `"types": ["jest", "node"]` to `compilerOptions` in `tsconfig.json`
so test globals (`describe`, `it`, `expect`) are visible without an explicit import.

---

#### Python — `SailPoint` module object in top-level `__init__.py`

**Namespace naming convention — strip the version from the class name:**

The `SailPoint` namespace uses resource names (`AccountsApi`), not versioned names
(`AccountsV1Api`).  Method names carry the version suffix (e.g. `list_accounts_v1_with_http_info`,
and future `list_accounts_v2_with_http_info`).  The import itself never changes when new
versions land.

**Usage:**

```python
# Preferred — flat imports from root package
from sailpoint import ApiClient, AccountsApi, SourcesApi, Configuration

cfg = Configuration()
api_client = ApiClient(cfg)                              # shared, works with any partition
accounts = AccountsApi(api_client).list_accounts_v1_with_http_info()
sources  = SourcesApi(api_client).list_sources_v1_with_http_info()

# Alternative — namespace style
from sailpoint import SailPoint, ApiClient, Configuration
api = SailPoint.AccountsApi(ApiClient(Configuration()))
api.list_accounts_v1_with_http_info()
```

**`generateInitPy()` generates `sailpoint/__init__.py`:**

```python
# Auto-generated by build-versioned-sdk.js — DO NOT EDIT
from sailpoint.api_client import ApiClient   # root-level shared client (see Task 15)
from sailpoint.accounts_v1.api.accounts_v1_api import AccountsV1Api as _AccountsV1Api
from sailpoint.transforms_v1.api.transforms_v1_api import TransformsV1Api as _TransformsV1Api

# Resource-named top-level exports (no version in the name)
AccountsApi   = _AccountsV1Api
TransformsApi = _TransformsV1Api

class SailPoint:
    AccountsApi   = _AccountsV1Api
    TransformsApi = _TransformsV1Api

from sailpoint.configuration import Configuration, ConfigurationParams
from sailpoint.paginator import Paginator
```

> **⚠️ Do NOT add versioned top-level exports** (e.g. `AccountsV1Api = _AccountsV1Api`)
> on a first release — there is nothing to be backward-compatible with.  Leave them out
> entirely.  Only add them if you are migrating a live SDK where users have existing code
> to preserve.

**Multi-version** (when `accounts_v1` and `accounts_v2` both exist): Python multiple
inheritance provides the cleanest combination — latest version first so its MRO wins on
name conflicts, older-version methods remain accessible:

```python
class _AccountsApiCombined(_AccountsV2Api, _AccountsV1Api): pass

class SailPoint:
    AccountsApi = _AccountsApiCombined  # exposes both v1 and v2 methods
```

`generateInitPy()` in `build-versioned-sdk.js` handles this automatically.

For direct partition access (advanced users):

```python
from sailpoint.accounts_v1 import AccountsV1Api   # direct partition import
```

**Class-detection regex pitfall:**

When scanning generated `*_api.py` files to find the API class name, use the pattern
`/^class (\w+)[(:]/m` — **not** `/^class (\w+)\(/m`.  The Python generator sometimes emits
classes without a base class (`class AccountsV1Api:`) and the parenthesis-only pattern
silently skips them, causing the build to report "No API classes found" and skip `__init__.py`
regeneration entirely.

**Method name suffix pitfall:**

The Python generator appends the version suffix to every method.  The generated method is
`list_accounts_v1_with_http_info`, **not** `list_accounts_with_http_info`.  All tests and
example code must use the suffixed names.

**`paginator.py` update:**

`sailpoint/paginator.py` contains hardcoded imports from the old monolithic `sailpoint.v3`
package.  After migration, update it to import from the new partition:

```python
# OLD
from sailpoint.v3.api.search_api import SearchApi
from sailpoint.v3.models.search import Search

# NEW
from sailpoint.search_v1.api.search_v1_api import SearchV1Api as SearchApi
from sailpoint.search_v1.models.search import Search
# also update method calls: search_post → search_post_v1
```

---

#### PowerShell — barrel module `PSSailpoint.psm1`

After the build generates `PSSailpoint/{partition}_v1/` modules, auto-generate a root barrel
module that dot-sources or imports each partition module. PowerShell users can then import a
single module to get access to all partitions:

```powershell
# Auto-generated by build-versioned-sdk.js — DO NOT EDIT
Import-Module -Name '.\PSSailpoint\accounts_v1\src\PSSailpoint.AccountsV1' -Force
Import-Module -Name '.\PSSailpoint\transforms_v1\src\PSSailpoint.TransformsV1' -Force
# ...
```

Users load the single barrel:

```powershell
Import-Module -Name '.\PSSailpoint\PSSailpoint' -Force
# All partition cmdlets are now available
Get-AccountsV1   # from accounts_v1
Get-TransformV1  # from transforms_v1
```

---

#### Go — auto-generated `client.go` _(already implemented)_

The `generateClientGo()` function in `build-versioned-sdk.js` produces `client.go` which exposes
each partition as a typed field on `APIClient`:

```go
apiClient.Accounts.AccountsAPI.ListAccountsV1(ctx).Execute()
apiClient.Transforms.TransformsAPI.ListTransformsV1(ctx).Execute()
```

This is the Go equivalent of the `SailPoint` namespace — a single `APIClient` gives access to
all partitions without additional imports.

---

### 17. Verify the Developer Portal Reference Path Convention _(was Task 15)_

The new docs workflow syncs to `docs/tools/sdk/go/Reference/{resource}` (e.g.,
`Reference/api_accounts_v1/Methods`), using the full partition directory name as the path
segment. Confirm with the developer portal team that this naming convention matches the expected
URL structure, and update the `rsync` destination paths if needed.

---

## Summary of File Changes

| Action  | File / Directory                              | Notes                                          |
|---------|-----------------------------------------------|------------------------------------------------|
| Added   | `sdk-resources/build-versioned-sdk.js`        | Core build script                              |
| Added   | `sdk-resources/sync-go-modules.js`            | Go workspace manager (Go-specific)             |
| Added   | `api_{resource}_v1/` (×N)                     | Per-partition generated packages               |
| Added   | `go.work`, `go.work.sum`                      | Go multi-module workspace (Go-specific)        |
| Removed | `api_v3/`, `api_beta/`, `api_v2024/`, `api_v2025/`, `api_v2026/` | Old monolithic packages   |
| Removed | `sdk-resources/v3-config.yaml`                | Config now generated dynamically               |
| Removed | `sdk-resources/beta-config.yaml`              | Config now generated dynamically               |
| Removed | `sdk-resources/v2024-config.yaml`             | Config now generated dynamically               |
| Removed | `sdk-resources/v2025-config.yaml`             | Config now generated dynamically               |
| Removed | `sdk-resources/v2026-config.yaml`             | Config now generated dynamically               |
| Removed | `sdk-resources/prescript.js`                  | Logic moved inline to build script             |
| Modified | `.github/actions/build-sdk/action.yaml`      | Replace version steps with versioned build     |
| Modified | `.github/workflows/push_sdk_docs_to_dev_portal.yaml` | Dynamic loop over partitions          |
| Modified | `Makefile`                                    | Use new build script; add partition target     |
| Modified | `client.go`                                   | Now auto-generated; do not edit manually       |
| Modified | `validation_test.go`                          | Update imports and API call patterns           |
| Modified | `go.mod`                                      | Simplify; remove old version deps (Go-specific)|
| Modified | `.gitignore`                                  | Add `.sdk-build-tmp/` and `build-errors/`      |

---

## Checklist for Applying to Another SDK

- [ ] Create `sdk-resources/build-versioned-sdk.js` (adapt config generation for target language)
- [ ] Move prescript logic inline (operate on temp copy, not source)
- [ ] Create language-equivalent of `sync-go-modules.js` if needed
- [ ] Delete old monolithic version packages (`v3`, `beta`, `v2024`, `v2025`, `v2026`)
- [ ] Delete old static version config YAMLs
- [ ] Delete standalone `prescript.js` (if present)
- [ ] Update `build-sdk` GitHub Action (replace version steps, add workspace wiring, fix test command)
- [ ] Update docs-push workflow (dynamic partition loop)
- [ ] Update `Makefile` / equivalent (use new script, add `build-partition` target)
- [ ] Auto-generate top-level client/index/barrel file from discovered partitions
- [ ] Add grouped-access entry point (`SailPoint` namespace / module object / barrel — see Task 16)
- [ ] **Migrate tests to V1 equivalents — do NOT delete them** (see Task 11 warning)
- [ ] Update `go.mod` / `requirements.txt` / `package.json` root dependency file
- [ ] Commit generated workspace file (`go.work` or equivalent)
- [ ] Update `.gitignore` (add `.sdk-build-tmp/`, `build-errors/`)
- [ ] Confirm developer portal reference path convention with portal team

### Python-specific checklist additions

- [ ] Class-detection regex uses `/^class (\w+)[(:]/m` (not `\(` alone) to catch classes without
      a base class
- [ ] `generateInitPy()` copies + patches `exceptions.py`, `api_response.py`, `rest.py`,
      `api_client.py` into `sailpoint/` root — not a re-export wrapper (see Task 15)
- [ ] `api_client.py` patch: remove hardcoded `import sailpoint.{partition}.models` and replace
      the `getattr(sailpoint.{partition}.models, klass)` lookup with `sys.modules` dynamic search
- [ ] `api_client.py` patch: fix `rest.py` and `api_client.py` imports of `exceptions`,
      `api_response`, and `rest` to reference `sailpoint.*` not `sailpoint.{partition}.*`
- [ ] `sailpoint/__init__.py` does **not** include versioned top-level exports (`AccountsV1Api`)
      on a first release — only resource-named exports (`AccountsApi`)
- [ ] All test method calls use the `_v1` suffix (e.g. `list_accounts_v1_with_http_info`) —
      the generator appends the version to every method name
- [ ] `sailpoint/paginator.py` updated: imports from `sailpoint.search_v1`, method calls changed
      from `search_post` → `search_post_v1` (and `_with_http_info` variant)
- [ ] Tests import `ApiClient` from `sailpoint` root, not from a specific partition sub-package
