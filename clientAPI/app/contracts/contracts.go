package contracts

import (
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
