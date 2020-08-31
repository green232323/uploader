package main

import (
	"context"
	"github.com/dnahurnyi/uploader/portDomain/app"
	"github.com/dnahurnyi/uploader/portDomain/app/storage"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"

	"google.golang.org/grpc"
)

const (
	dbURL = "DB_URL"
	port  = "PORT"
)

func main() {
	ownPort := os.Getenv(port)
	if len(ownPort) == 0 {
		log.Fatalf("failed to get own port from env")
	}

	lis, err := net.Listen("tcp", ":"+ownPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databaseURL := os.Getenv(dbURL)
	if len(databaseURL) == 0 {
		log.Fatalf("failed to get database credentials from env")
	}

	portRepo, err := storage.NewMongoRepository(ctx, databaseURL)
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
