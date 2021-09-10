package main

import (
	"context"
	"log"

	"github.com/speedrunsh/portal"
)

func (s *server) RunCommand(ctx context.Context, in *portal.Command) (*portal.Response, error) {
	log.Printf("Received command:%s", in.GetName())
	return &portal.Response{Content: "ran " + in.GetName()}, nil
}
