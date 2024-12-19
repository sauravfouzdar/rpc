package main

import (
	"net"
	"log"
	"context"

	pb "rpc/rpc/protos/sum"

	"google.golang.org/grpc"
)


// server implements the SumService interface
type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) Add(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: req.A + req.B}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})

	log.Println("Starting server on port :1234")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}