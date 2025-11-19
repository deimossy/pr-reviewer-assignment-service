.PHONY: up down logs build clean test bench create-network send-orders codegen logger

bin/codegen_install:
	@echo "Installing codegen tool..."
	@mkdir -p bin
	@GOBIN=$(PWD)/bin go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

codegen: bin/codegen_install
	bin/oapi-codegen --package=gen --generate types docs/openapi.yaml > internal/dto/gen/types.gen.go

test:
	go test -v ./...

#bench:
#	go test ./internal/services/example -bench=. -benchmem

up:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down

logs:
	docker compose logs -f

logger:
	docker compose logs -f pr_reviewer_assignment_service

clean:
	docker compose down -v