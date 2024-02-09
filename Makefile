.PHONY: specs
specs:	
	git clone https://github.com/sailpoint-oss/api-specs.git api-specs

.PHONY: clean-specs
clean-specs:	
	rm -rf ./api-specs

.PHONY: build
build:	
	node sdk-resources/prescript.js api-specs/
	rm -rf ./api_v3
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v3.yaml -g go -o api_v3 --global-property skipFormModel=false --config sdk-resources/v3-config.yaml 
	node sdk-resources/postscript.js ./api_v3
	rm -rf ./api_beta
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.beta.yaml -g go -o api_beta --global-property skipFormModel=false --config sdk-resources/beta-config.yaml
	node sdk-resources/postscript.js ./api_beta

.PHONY: test
test:	
	go get -d ./...
	go install
	go test
	