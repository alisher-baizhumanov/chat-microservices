LOCAL_BIN:=$(CURDIR)/bin

install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

prepare:
	make tidy
	make test
	make build
	make lint

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./pkg/... ./services/auth/... ./services/chat-server/... --config .golangci.pipeline.yaml

test:
	go test -v ./pkg/... ./services/auth/... ./services/chat-server/...

build:
	go build -o ./bin/auth -mod vendor -v ./services/auth/
	go build -o ./bin/chat-server -mod vendor -v ./services/chat-server/

tidy:
	cd pkg && go mod tidy && cd .. && \
	cd protos && go mod tidy && cd .. && \
	cd services/auth && go mod tidy -e && cd ../.. && \
	cd services/chat-server && go mod tidy -e && cd ../..
	go work sync && go work vendor

generate:
	make generate-user-api
	make generate-chat-api

generate-user-api:
	mkdir -p protos/generated/user-v1
	protoc --proto_path protos/sources/user-v1 \
    	--go_out=protos/generated/user-v1 --go_opt=paths=source_relative \
    	--plugin=protoc-gen-go=bin/protoc-gen-go \
    	--go-grpc_out=protos/generated/user-v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc protos/sources/user-v1/user.proto

generate-chat-api:
	mkdir -p protos/generated/chat-v1
	protoc --proto_path protos/sources/chat-v1 \
    	--go_out=protos/generated/chat-v1 --go_opt=paths=source_relative \
    	--plugin=protoc-gen-go=bin/protoc-gen-go \
    	--go-grpc_out=protos/generated/chat-v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc protos/sources/chat-v1/chat.proto
