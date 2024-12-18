package main

import (
	"net/rpc"
	"net/http"
	"net"
	"log"
	"time"
)


type Args struct {}

type TimeServer int64

func main() {
	// Create a new RPC server
	timeserver := new(TimeServer)
	rpc.Register(timeserver)

	// Register the RPC server as an HTTP handler
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	defer listener.Close()

	log.Printf("Serving RPC server on port 1234")

	http.Serve(listener, nil)

}


func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = int64(time.Now().Unix())
	return nil
}