package main

import (
	"trpc.group/trpc-go/tnet/log"
	"trpc.group/trpc-go/trpc-gateway/core/config"
	"trpc.group/trpc-go/trpc-gateway/core/service/fhttp"
)

func main() {
	s := config.NewServer()

	fhttp.RegisterFastHTTPService(s.Service("trpc.http.service"))
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
