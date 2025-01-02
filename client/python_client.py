# write a client that sends rpc request to the server and prints the response

import grpc
import sum_pb2
import sum_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:1234')
    stub = sum_pb2_grpc.SumServiceStub(channel)
    response = stub.Sum(sum_pb2.SumRequest(a=1, b=2))
    print(response.result)


if __name__ == "__main__":
    run()
