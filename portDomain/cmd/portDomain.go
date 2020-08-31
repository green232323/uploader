package main

import (
	"github.com/dnahurnyi/uploader/portDomain/app"
	"log"
	"net"

	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPortDomainServer(grpcServer, &app.Server{})

	log.Println("service started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
