package main

import (
	"trpc.group/trpc-go/trpc-go"
)

type UserServiceImpl struct{}

func main() {
	s := trpc.NewServer()
	s.Serve()
}
