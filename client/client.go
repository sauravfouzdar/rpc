package main

import (
	"log"
	"net/http"
	"net/rpc"
	"fmt"
)

type httpHandler struct {
	rpcClient *rpc.Client
}

type Args struct{}

func (h *httpHandler) GetTime(w http.ResponseWriter, r *http.Request) {
	// Prepare arguments and response
	var reply int64
	args := Args{}

	// Call the remote procedure
	err := h.rpcClient.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		http.Error(w, "Error calling remote procedure: "+err.Error(), http.StatusInternalServerError)
		log.Println("Error in TimeServer.GiveServerTime: ", err)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"server_time": ` + fmt.Sprintf("%d", reply) + `}`))
	log.Printf("Server time received: %d", reply)
}

func initRPCClient() (*rpc.Client, error) {
	// Create a new RPC client
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:1234")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func initRoutes(rpcClient *rpc.Client) *http.ServeMux {
	router := http.NewServeMux()
	h := &httpHandler{rpcClient: rpcClient}
	router.HandleFunc("/time", h.GetTime) // Pass the function, not the result of calling it
	return router
}

func main() {
	// Initialize RPC client
	rpcClient, err := initRPCClient()
	if err != nil {
		log.Fatal("Failed to connect to RPC server: ", err)
	}
	defer rpcClient.Close()

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: initRoutes(rpcClient),
	}

	// Log and start the server
	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
