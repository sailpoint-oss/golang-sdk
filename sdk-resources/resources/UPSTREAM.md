# Upstream template base

These templates are forks of the stock OpenAPI Generator Go templates from
**openapi-generator 7.18.0** (`go/*.mustache` inside the jar).

`developerSite_code_examples`, `docs_methods_index`, and `docs_models_index`
are SailPoint-only templates with no upstream counterpart.

## Upgrading the generator

Extract the stock templates for the current base and the new version, then
3-way merge each shared template so local changes are kept:

```sh
unzip -q openapi-generator-cli-<old>.jar 'go/*' -d stock-old
unzip -q openapi-generator-cli-<new>.jar 'go/*' -d stock-new
for f in stock-old/go/*.mustache; do
  f=$(basename "$f")
  git merge-file sdk-resources/resources/$f stock-old/go/$f stock-new/go/$f
done
```

Resolve any conflicts, then update the version in this file,
`.github/actions/build-sdk/action.yaml`, and `sdk-resources/build-versioned-sdk.js`.

Codegen changes are not visible in the templates. Before merging an upgrade,
regenerate with both jars against the same `api-specs` checkout and diff the
output for removed models or changed field types. For example, 7.18 stopped
creating inline models for annotation-only `allOf` members. That was
countered with `typeAnnotationOnlyAllOfMembers` in `build-versioned-sdk.js`.
