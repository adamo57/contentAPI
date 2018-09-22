package main

import (
	"context"
	"log"

	"github.com/adamo57/contentAPI/api"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	// connect to the gRPC server.
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect to server")
	}
	defer conn.Close()

	// register as a new client to Post Service.
	c := api.NewPostClient(conn)

	// make request to the GetPost route on the Post Service.
	resp, err := c.GetPost(context.Background(), &api.PostSlug{Name: "Test Post"})
	if err != nil {
		log.Fatal("error when calling GetPost")
	}

	log.Print(resp)
}
