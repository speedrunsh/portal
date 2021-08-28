package main

import (
	"context"
	"log"
	"net"

	"github.com/speedrunsh/portal/command"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:1337"

type server struct {
	command.UnimplementedPortalServer
}

// func (s *server) Echo(ctx context.Context, in *service.Empty) (*service.Empty, error) {
// 	log.Printf("Received ping")
// 	return &service.Empty{}, nil
// }

// func (s *server) ServiceRestart(ctx context.Context, in *service.Service) (*service.Response, error) {
// 	conn, err := dbus.NewWithContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()
// 	responseChan := make(chan string, 1)
// 	serviceName := fmt.Sprintf("%s.service", in.GetName())
// 	_, err = conn.RestartUnitContext(ctx, serviceName, "replace", responseChan)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := <-responseChan
// 	return &service.Response{Content: res}, nil

// }

func (s *server) RunCommand(ctx context.Context, in *command.Command) (*command.Response, error) {
	log.Printf("Received command")
	return &command.Response{Content: "command done"}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	command.RegisterPortalServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
