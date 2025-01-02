## RPC 
` A simple grpc server with a simple client that sends a request to the server and the server responds with a message.`


## Generate the proto file

# python
```bash
python -m pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. calculator.proto
```
# go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=. calculator.proto
```

## How to run
1. Clone the repository
2. Run the server - `go run rpc/rpc_server.go`
3. Run the client - `go run rpc/client.go`

