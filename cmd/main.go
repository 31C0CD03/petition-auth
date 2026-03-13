package main

import (
	"log"
	"net"

	cfg "github.com/31c0cd03/petition-pkg/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := cfg.LoadServiceConfig()
	lis, err := net.Listen("tcp", cfg.Port)

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	log.Printf("listening on %v\n", cfg.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
