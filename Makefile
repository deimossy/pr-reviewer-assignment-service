bin/codegen_install:
	@echo "Installing codegen tool..."
	@mkdir -p bin
	@GOBIN=$(PWD)/bin go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

codegen: bin/codegen_install
	bin/oapi-codegen --package=gen --generate types docs/openapi.yaml > internal/dto/gen/types.gen.go

