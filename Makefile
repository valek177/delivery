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

generate-geo-client:
	@rm -rf internal/generated/clients/geosrv
	@protoc --go_out=internal/generated/clients --go-grpc_out=internal/generated/clients api/proto/geo_service.proto
