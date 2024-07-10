module github.com/alisher-baizhumanov/chat-microservices/services/auth

go 1.22

toolchain go1.22.4

require (
	github.com/alisher-baizhumanov/chat-microservices/pkg v0.0.0
	github.com/alisher-baizhumanov/chat-microservices/protos v0.0.0
	github.com/brianvoe/gofakeit/v7 v7.0.4
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)

replace (
	github.com/alisher-baizhumanov/chat-microservices/pkg v0.0.0 => ../../pkg
	github.com/alisher-baizhumanov/chat-microservices/protos v0.0.0 => ../../protos
)
