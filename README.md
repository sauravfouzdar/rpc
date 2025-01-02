## RPC 
A simple grpc server with a client that sends a request to the server and the server responds with a message(result).


## Generate the proto file

# python
```bash
python -m pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. sum.proto
```
# go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=. sum.proto
```

## How to run
1. Install go
2. Install the dependencies - `go mod tidy`
3. Install python and the grpcio-tools package
4. Generate the proto file
5. Run the server - `go run rpc/rpc_server.go`
6. Run the client - `go run rpc/client.go`



