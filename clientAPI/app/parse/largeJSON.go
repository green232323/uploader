package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnahurnyi/uploader/clientAPI/app/contracts"
	"github.com/dnahurnyi/uploader/clientAPI/app/models"
	"io"
)

type largeJSONParser struct{}

func (p largeJSONParser) Parse(stream io.ReadCloser) (func() (contracts.Storable, error), error) {
	dec := json.NewDecoder(stream)
	// Needed for this specific type of JSON, where real data starts inside of the object
	if err := passToken(dec, '{'); err != nil {
		return nil, err
	}
	return func() (contracts.Storable, error) {
		return p.parseJSONObject(dec)
	}, nil
}

func (p largeJSONParser) parseJSONObject(dec *json.Decoder) (contracts.Storable, error) {
	if !dec.More() {
		return nil, nil
	}
	key, err := dec.Token()
	if err != nil {
		return nil, err
	}
	port := models.Port{}
	err = dec.Decode(&port)
	if err != nil {
		return nil, err
	}
	keyStr, ok := key.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("failed to parse key: %v", key))
	}
	port.Key = keyStr
	return port, nil
}

func passToken(dec *json.Decoder, token json.Delim) error {
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if delim, ok := t.(json.Delim); !ok || delim != token {
		return errors.New("unexpected token: " + string(token))
	}
	return nil
}

func LargeJsonParser() contracts.Parser {
	return largeJSONParser{}
}
