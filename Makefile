.PHONY: generate
generate:
	docker run --rm -v "$$(pwd):/local" -e JAVA_OPTS="-Xmx1024M -Xss32M" openapitools/openapi-generator-cli:v7.4.0 generate \
		-i /local/api/swagger.json \
		-g go \
		-o /local/api \
		--package-name api \
		--additional-properties=enumClassPrefix=true \
		--global-property apiTests=false,modelTests=false,apiDocs=false,modelDocs=false
	go run ./tools/fix-openapi-gen/main.go
	go run ./tools/gen-execute-all/main.go
	goimports -w ./api
