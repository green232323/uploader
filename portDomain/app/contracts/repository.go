package contracts

import (
	"context"
	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"
)

type PortRepository interface {
	SaveOrUpdate(key string, value *pb.Port) error
	Get(ctx context.Context, key string) (*pb.Port, error)
}
