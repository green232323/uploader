package main

import (
	"context"
	"github.com/dnahurnyi/uploader/portDomain/app"
	"github.com/dnahurnyi/uploader/portDomain/app/storage"
	"log"
	"net"
	"time"

	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	portRepo, err := storage.NewMongoRepository(ctx, "mongo://localhost:27017")
	if err != nil {
		log.Fatalf("failed to connect to database, err: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPortDomainServer(grpcServer, &app.Server{DB: portRepo})

	log.Println("service started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
