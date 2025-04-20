package yaml

import (
	"github.com/adnpa/fern/pkg/encoding"
	"gopkg.in/yaml.v3"
)

const Name = "yaml"

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (codec) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
