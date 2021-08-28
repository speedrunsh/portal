.PHONY: requirements command

all: requirements command

requirements:
	@go mod download
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

command:
	protoc --go_out=paths=source_relative:.. --go-grpc_out=paths=source_relative:.. command/command.proto