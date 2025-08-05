package main

import (
	"flag"

	"trpc.group/trpc-go/trpc-gateway/core/config"
	"trpc.group/trpc-go/trpc-gateway/core/service/fhttp"
	"trpc.group/trpc-go/trpc-go/log"

	// Register file loader
	_ "trpc.group/trpc-go/trpc-gateway/core/loader/file"

	// Register upstream protocol
	_ "trpc.group/trpc-go/trpc-gateway/core/service/protocol/fasthttp"
	_ "trpc.group/trpc-go/trpc-gateway/core/service/protocol/grpc"
	_ "trpc.group/trpc-go/trpc-gateway/core/service/protocol/http"
	_ "trpc.group/trpc-go/trpc-gateway/core/service/protocol/trpc"
)

func main() {
	flag.Parse()
	s := config.NewServer()

	fhttp.RegisterFastHTTPService(s.Service("trpc.http.service"))
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
