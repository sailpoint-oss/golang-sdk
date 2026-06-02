.PHONY: specs
specs:
	git clone https://github.com/sailpoint-oss/api-specs.git api-specs

.PHONY: clean-specs
clean-specs:
	rm -rf ./api-specs

# Path to the apis/ directory inside the checked-out api-specs repo.
# Override with: make build APIS_DIR=/path/to/apis
APIS_DIR ?= api-specs/idn/src/main/yaml/apis

# Build all per-partition versioned SDK packages.
# New partitions are auto-discovered — no Makefile changes needed when endpoints are added.
.PHONY: build
build:
	node sdk-resources/build-versioned-sdk.js $(APIS_DIR)
	rm -rf ./api_generic
	java -jar openapi-generator-cli.jar generate -i sdk-resources/generic-api.yaml -g go -o api_generic --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/generic-config.yaml
	node sdk-resources/postscript.js ./api_generic
	rm -rf ./api_nerm
	java -jar openapi-generator-cli.jar generate -i api-specs/nerm/openapi.yaml -g go -o api_nerm --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/nerm-config.yaml
	node sdk-resources/postscript.js ./api_nerm
	rm -rf ./api_nerm_v2025
	java -jar openapi-generator-cli.jar generate -i api-specs/nerm/v2025/v2025.yaml -g go -o api_nerm_v2025 --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/v2025-nerm-config.yaml
	node sdk-resources/postscript.js ./api_nerm_v2025

# Build a single partition (fast iteration during development).
# Usage: make build-partition PARTITION=entitlements
.PHONY: build-partition
build-partition:
	node sdk-resources/build-versioned-sdk.js $(APIS_DIR) --partition $(PARTITION)

.PHONY: sync-modules
sync-modules:
	node sdk-resources/sync-go-modules.js

.PHONY: test
test: sync-modules
	go build -o /dev/null .
	go test ./...
