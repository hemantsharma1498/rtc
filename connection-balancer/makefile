# This is responsible for the following:
#       - compiling golang packages
#       - compiling proto files
#
#       TODOS
#       1. Pkg/proto
#               - to not exist
#               - input and output both in ./proto/
#       2.
#
#

#TOOLS
PROTO_SRC := ./proto
PROTO_OUT := ./pkg/proto
GO_MAIN := ./cmd/connection-balancer/main.go
GO_OUT := ./bin
PROTOC := protoc 
# Ensuring tools are installed
.PHONY: tools
tools:
		go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
		go get google.golang.org/grpc
		go mod tidy

.PHONY: proto
proto:
	mkdir -p $(PROTO_OUT)
	$(PROTOC) --proto_path=$(PROTO_SRC) --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_SRC)/*.proto

# Building the application
.PHONY: build
build: proto
		mkdir -p $(GO_OUT)
		go build -o $(GO_OUT) $(GO_MAIN)

# Running the application
.PHONY: run
run: build
	$(GO_OUT)/main

