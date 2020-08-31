package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dnahurnyi/uploader/clientAPI/app/contracts"
	"github.com/dnahurnyi/uploader/clientAPI/app/models"
	"google.golang.org/grpc"
	"io"
	pb "portDomain/proto/github.com/dnahurnyi/uploader/portDomain"
)

// PortDomain is a client for data delivery to PortDomainService
type portDomain struct {
	url    string
	client pb.PortDomainClient
}

func PortDomainClient(url string) (contracts.DomainClient, error) {
	pd := &portDomain{
		url: url,
	}
	err := pd.Connect()
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (pd *portDomain) Connect() error {
	conn, err := grpc.Dial(pd.url, grpc.WithInsecure())
	if err != nil {
		return errors.New("did not connect to port domain, err: " + err.Error())
	}
	pd.client = pb.NewPortDomainClient(conn)
	return nil
}

// Receiver receive data and sends it to another service
func (pd *portDomain) Send(ctx context.Context) (func(contracts.Storable, bool) error, error) {
	stream, err := pd.client.LoadPorts(ctx)
	if err != nil {
		return nil, err
	}
	closeF := func() error {
		_, err = stream.CloseAndRecv()
		return err
	}
	return func(port contracts.Storable, close bool) error {
		if close {
			return closeF()
		}
		protoPort := &pb.Port{}
		err := json.Unmarshal(port.Present(), protoPort)
		if err != nil {
			return err
		}
		err = stream.Send(protoPort)
		if err != nil {
			if err == io.EOF {
				return closeF()
			}
		}
		return err
	}, nil
}

func (pd *portDomain) Get(ctx context.Context, key string) (contracts.Storable, error) {
	port, err := pd.client.GetPortByID(ctx, &pb.PortID{
		Key: key,
	})
	if err != nil {
		return nil, err
	}
	bPort, err := json.Marshal(port)
	portModel := models.Port{}
	err = json.Unmarshal(bPort, &portModel)
	if err != nil {
		return nil, err
	}
	return portModel, nil
}
