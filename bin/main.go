package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/coreos/go-systemd/v22/dbus"
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

func (s *server) ServiceRestart(ctx context.Context, in *portal.Service) (*portal.Response, error) {
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	responseChan := make(chan string, 1)
	serviceName := fmt.Sprintf("%s.service", in.GetName())
	_, err = conn.RestartUnitContext(ctx, serviceName, "replace", responseChan)
	if err != nil {
		return nil, err
	}

	res := <-responseChan
	return &portal.Response{Content: res}, nil

}

func (s *server) RunCommand(ctx context.Context, in *portal.Command) (*portal.Response, error) {
	log.Printf("Received command:%s", in.GetName())
	return &portal.Response{Content: "ran " + in.GetName()}, nil
}

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
