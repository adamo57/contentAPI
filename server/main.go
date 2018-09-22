package main

import (
	"log"
	"net"

	"github.com/adamo57/contentAPI/api"
	"google.golang.org/grpc"
)

func main() {
	// listener listens on certain port.
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("failed to listen")
	}

	// new api server instance.
	s := api.Server{}

	grpcServer := grpc.NewServer()

	// make connection to the post service.
	api.RegisterPostServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
