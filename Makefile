.PHONY: specs
specs:
	git clone https://github.com/sailpoint-oss/api-specs.git api-specs

.PHONY: clean-specs
clean-specs:
	rm -rf ./api-specs

.PHONY: build
build:
	node sdk-resources/prescript.js api-specs/idn
	rm -rf ./api_v3
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v3.yaml -g go -o api_v3 --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/v3-config.yaml
	node sdk-resources/postscript.js ./api_v3
	rm -rf ./api_beta
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.beta.yaml -g go -o api_beta --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/beta-config.yaml
	node sdk-resources/postscript.js ./api_beta
	rm -rf ./api_v2024
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v2024.yaml -g go -o api_v2024 --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/v2024-config.yaml
	node sdk-resources/postscript.js ./api_v2024
	rm -rf ./api_v2025
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v2025.yaml -g go -o api_v2025 --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/v2025-config.yaml
	node sdk-resources/postscript.js ./api_v2025
	rm -rf ./api_generic
	java -jar openapi-generator-cli.jar generate -i sdk-resources/generic-api.yaml -g go -o api_generic --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/generic-config.yaml
	node sdk-resources/postscript.js ./api_generic
	rm -rf ./api_v2026
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v2026.yaml -g go -o api_v2026 --global-property skipFormModel=false,apiDocs=true,modelDocs=true --config sdk-resources/v2026-config.yaml
	node sdk-resources/postscript.js ./api_v2026

.PHONY: test
test:
	go get -d ./...
	go install
	go test
