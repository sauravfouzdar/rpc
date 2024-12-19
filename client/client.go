package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	pb "rpc/rpc/protos/sum"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RequestPayload struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponsePayload struct {
	Result int `json:"result"`
}

// gRPC client used to call the gRPC server.
//var grpcClient pb.SumServiceClient

func addHandler(w http.ResponseWriter, r *http.Request) {

	var reqPayload RequestPayload

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Make the Add RPC call
	conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewSumServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SumRequest{A: int32(reqPayload.A), B: int32(reqPayload.B)}
	res, err := grpcClient.Add(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}

	fmt.Printf("Result of %d + %d = %d\n", req.A, req.B, res.Result)

	// Create and send the response
	ResponsePayload := ResponsePayload{Result: int(res.Result)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload)

}

func main() {

	// Set up a connection to the server.
	r := mux.NewRouter()

	r.HandleFunc("/add", addHandler).Methods("POST")

	fmt.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", r)

}
