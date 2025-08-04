package main

import (
	helloworld "github.com/kickastone/trpc-demo/proto/upstream"
	"trpc.group/trpc-go/trpc-go"
)

type HelloworldServiceImpl struct{}

func main() {
	s := trpc.NewServer()
	helloworld.RegisterAuthService(s, &HelloworldServiceImpl{})
	s.Serve()
}
