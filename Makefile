.PHONY: specs
specs:	
	git clone https://github.com/sailpoint-oss/api-specs.git api-specs

.PHONY: clean-specs
clean-specs:	
	rm -rf ./api-specs

.PHONY: build
build:	
	node sdk-resources/prescript.js api-specs/
	rm -rf ./v3
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.v3.yaml -g go -o v3 --global-property skipFormModel=false --global-property apiTests=false --config sdk-resources/v3-config.yaml
	node sdk-resources/postscript.js ./v3
	rm -rf ./beta
	java -jar openapi-generator-cli.jar generate -i api-specs/idn/sailpoint-api.beta.yaml -g go -o beta --global-property skipFormModel=false --global-property apiTests=false --config sdk-resources/beta-config.yaml
	node sdk-resources/postscript.js ./beta

test:	
	go get -d ./...
	go install
	go test
	