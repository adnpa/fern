package proto

import (
	"github.com/adnpa/fern/pkg/encoding"
	"google.golang.org/protobuf/proto"
)

const Name = "proto"

type codec struct{}

func (c codec) Marshal(v any) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (c codec) Unmarshal(data []byte, v any) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (c codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterCodec(codec{})
}
