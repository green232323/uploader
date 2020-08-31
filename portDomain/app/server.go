package app

import (
	"context"
	"fmt"
	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"
	"io"
)

type Server struct {
}

func (s *Server) LoadPorts(stream pb.PortDomain_LoadPortsServer) error {
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}
		if err != nil {
			fmt.Printf("client died, err: %v\n", err)
			return err
		}
		fmt.Println("Got port: ", port)
	}
}

func (s *Server) Get(ctx context.Context, portID *pb.PortID) (*pb.Port, error) {
	fmt.Println("Return port by id: ", portID)
	return nil, nil
}
