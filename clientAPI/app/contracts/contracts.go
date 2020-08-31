package contracts

import (
	"context"
	"io"
)

type Parser interface {
	// Parse gets a stream and return function to parse an infinite stream and preparation error
	Parse(stream io.ReadCloser) (func() (Storable, error), error)
}

// Storable is unified form of data transfer
type Storable interface {
	Present() []byte
}

type DomainClient interface {
	// Send returns function to send data to another service with stop option and preparation error
	Send(ctx context.Context) (func(Storable, bool) error, error)
	Get(ctx context.Context, key string) (Storable, error)
}
