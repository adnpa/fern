package json

import (
	"encoding/json"

	"github.com/adnpa/fern/pkg/encoding"
)

const Name = "json"

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (codec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
