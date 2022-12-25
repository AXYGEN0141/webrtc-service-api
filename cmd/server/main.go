package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()

	fmt.Println("Server is running on port:", port)
	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
