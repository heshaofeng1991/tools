.PHONY: tidy
tidy:
	@echo "go mod tidy..."
	@go mod tidy -compat=1.9

.PHONY: build
build:
	@echo "build admin server..."
	@go build main.go

.PHONY: server
server:
	@echo "start admin server..."
	@go run main.go -confFile ./config/config-dev.yaml