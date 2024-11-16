package main

import (
	"fmt"
	"json-parser/jsonhandler"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         "8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  20 * time.Second,
	}
	http.HandleFunc("/parse", jsonhandler.ParseJSONHandler)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(server.Serve(listener))

}
