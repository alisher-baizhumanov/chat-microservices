include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATIONS_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=55010 dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"
MIGRATION_NAME=""

install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.10
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

prepare:
	make tidy
	make lint
	make test
	make build

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./pkg/... ./services/auth/... ./services/chat-server/... --config .golangci.pipeline.yaml

test:
	go test -v ./pkg/... ./services/auth/... ./services/chat-server/...

build:
	go build -o ./bin/auth -mod vendor -v ./services/auth/
	go build -o ./bin/chat-server -mod vendor -v ./services/chat-server/

tidy:
	go mod tidy
	go mod vendor

mock:
	GOBIN=$(LOCAL_BIN) go generate ./...

generate:
	make generate-user-api
	make generate-chat-api

generate-user-api:
	mkdir -p protos/generated/user-v1
	protoc --proto_path api/proto/user-v1 \
		--proto_path=vendor.protogen \
    	--go_out=protos/generated/user-v1 --go_opt=paths=source_relative \
    	--plugin=protoc-gen-go=bin/protoc-gen-go \
    	--go-grpc_out=protos/generated/user-v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out lang=go:protos/generated/user-v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		--grpc-gateway_out=protos/generated/user-v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		api/proto/user-v1/user.proto

generate-chat-api:
	mkdir -p protos/generated/chat-v1
	protoc --proto_path api/proto/chat-v1 \
		--proto_path=vendor.protogen \
    	--go_out=protos/generated/chat-v1 --go_opt=paths=source_relative \
    	--plugin=protoc-gen-go=bin/protoc-gen-go \
    	--go-grpc_out=protos/generated/chat-v1 --go-grpc_opt=paths=source_relative \
    	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out lang=go:protos/generated/chat-v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		--grpc-gateway_out=protos/generated/chat-v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		api/proto/chat-v1/chat.proto

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi

up:
	make up-auth
	make up-chat

up-auth:
	docker-compose up auth --build --detach

up-chat:
	docker compose up chat-server --build --detach

stop:
	docker-compose stop

down:
	docker-compose down --remove-orphans --volumes

migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

migration-create:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} create ${MIGRATION_NAME} sql
