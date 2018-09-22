package api

import (
	"log"

	context "golang.org/x/net/context"
)

// Server is the gRPC server
type Server struct {
}

// GetPost fetches a post with the given postslug.
func (s *Server) GetPost(ctx context.Context, slug *PostSlug) (*PostObject, error) {
	// do db stuff
	log.Println("Hit GetPost")

	return &PostObject{
		Id:      1,
		Name:    "Test Post",
		Content: "This is a test post...",
	}, nil
}
