# Exemplo servidor e client gRPC em go

`protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb`

`go run ./cmd/server/server.go`

`go run ./cmd/client/client.go`

# Testar

https://github.com/ktr0731/evans

`evans -r repl --host localhost --port 50051`