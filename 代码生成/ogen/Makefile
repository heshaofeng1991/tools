.PHONY: tidy
tidy:
	@echo "go mod tidy..."
	@go mod tidy -compat=1.9

.PHONY: generate
generate:
	@echo "go generate..."
	@go generate ./...

.PHONY: server
server:
	@echo "start server..."
	@go run main.go