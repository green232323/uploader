package app

import (
	"context"
	"fmt"
	"github.com/dnahurnyi/uploader/portDomain/app/contracts"
	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"
	"io"
)

type Server struct {
	DB contracts.PortRepository
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

func (s *Server) GetPortByID(ctx context.Context, portID *pb.PortID) (*pb.Port, error) {
	fmt.Println("Return port by id: ", portID)
	return nil, nil
}
