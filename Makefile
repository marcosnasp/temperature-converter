generate:
    protoc --go_out=./proto/gen/go --go-grpc_out=./proto/gen/go proto/temperature.proto

run-server:
    go run cmd/server/main.go

run-client:
    go run cmd/client/main.go $(ARGS)
