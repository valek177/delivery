APP_NAME=delivery

.PHONY: build test
build: test ## Build application
	mkdir -p build
	go build -o build/${APP_NAME} cmd/app/main.go

test: ## Run tests
	go test ./...

server:
	oapi-codegen -config configs/server.cfg.yaml api/openapi.yaml
.PHONY: server
