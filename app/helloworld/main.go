package main

import (
	"context"

	helloworld "github.com/kickastone/trpc-demo/proto/upstream"
	"trpc.group/trpc-go/tnet/log"
	"trpc.group/trpc-go/trpc-go"
)

type HelloworldServiceImpl struct{}

func main() {
	s := trpc.NewServer()
	helloworld.RegisterHelloServiceService(s, &HelloworldServiceImpl{})
	s.Serve()
}

func (s *HelloworldServiceImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	log.Infof("SayHello: %v", req)
	return &helloworld.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}
