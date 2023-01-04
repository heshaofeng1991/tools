$(info $(SHELL))

GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install github.com/heshaofeng1991/protoc-gen-go-http@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest

.PHONY: admin-proto
admin-proto:
	protoc -I=. -I=${GOPATH}/pkg/mod \
 	--gogofaster_out=. --gogofaster_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --govalidators_out=. --govalidators_opt=paths=source_relative \
    --go-http_out=. --go-http_opt=paths=source_relative \
    proto/admin/admin.proto
