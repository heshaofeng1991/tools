#! /usr/bin/make
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "gen" invokes the "goa" tool to generate the examples source code
# - "build" compiles the example microservices and client CLIs
# - "clean" deletes the output of "build"
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
GO_FILES=$(shell find . -type f -name '*.go')
GOA:=$(shell goa version 2> /dev/null)

.PHONY: all
## all :default run all cmdline for server
all: check-goa goa ent tidy gofmt lint format compile serve

.PHONY: local
## local :run more fast exclude gofmt, lint
local: check-goa goa ent tidy format compile serve

.PHONY: build
## build :build server
build: check-goa goa ent tidy gofmt lint format compile

# Used for github actions
.PHONY: build-release
build-release:
	@echo "Building release..."
	make check-goa goa ent compile

.PHONY: reload
## reload :only build server & run server
reload: compile serve

.PHONY: check-goa
## check-goa :check goa is existed, no existed will install
check-goa:
ifdef GOA
	@echo $(GOA)
else
	go install goa.design/goa/v3/...@v3
	@echo $(GOA)
endif

.PHONY: goa
## goa :auto generate goa design code by design files
goa:
	@echo GENERATING CODE...
	@goa version
	@rm -rf "./gen"
	@goa gen combo/design -o "./"

.PHONY: ent
## ent :auto generate ent database orm code by schema files
ent:
	@echo "Generating ent code..."
	go generate ./ent

.PHONY: tidy
## tidy :download need mod && clean unused mod && upgrade mod
tidy:
	@echo TIDYING CODE...
	@go mod tidy -compat=1.19

.PHONY: format
## format :format project code
format:
	@echo FORMAT PROJECT CODE...
	@go fmt ./...

.PHONY: lint
## lint :go linters aggregator
lint:
	@echo LINTING CODE...
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	@if [ "`golangci-lint run ./... | tee /dev/stderr`" ]; then \
  		echo "^ - Lint errors!" && echo && exit 1; \
  	fi

.PHONY: gofmt
## gofmt :enforce a stricter format than gofmt, while being backwards compatible.
gofmt:
	@echo FOFUMPT PROJECT CODE...
	@go install mvdan.cc/gofumpt@latest
	@gofumpt -l -w .

.PHONY: compile
## compile :compile server code
compile:
	@echo BUILDING CODE...
	@go build ./cmd/combo && go build ./cmd/combo-cli

.PHONY: serve
## serve :start server
serve:
	@echo "Serving..."
	./combo

.PHONY: test
## test :unit test of server code
test:
	@echo TESTING...
	@go test ./.. > /dev/null

.PHONY: check-freshness
## check-freshness :auto generate goa design code by design files
check-freshness:
	@if [ "`git diff -- . ':(exclude)go.mod' ':(exclude)go.sum' ':(exclude)files/cmd/openapi/http.go' | wc -l`" -gt "0" ]; then \
           echo "[ERROR] generated code not in-sync with design:"; \
           echo; \
           git status -s; \
           git --no-pager diff; \
           echo; \
           exit 1; \
	fi

.PHONY: help
## help: prints all cmdline of help message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

