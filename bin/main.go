package main

import (
	"log"
	"net"

	"github.com/speedrunsh/portal"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:1337"

type server struct {
	portal.UnimplementedPortalServer
}

// func (s *server) Echo(ctx context.Context, in *portal.Empty) (*portal.Empty, error) {
// 	log.Printf("Received ping")
// 	return &portal.Empty{}, nil
// }

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	portal.RegisterPortalServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
