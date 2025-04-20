package encoding

import (
	"reflect"
	"testing"
)

type codec struct{}

func (c codec) Marshal(v any) ([]byte, error) {
	return nil, nil
}

func (c codec) Unmarshal(data []byte, v any) error {
	return nil
}

func (c codec) Name() string {
	return "test"
}

func TestRegisterGetCodec(t *testing.T) {
	type args struct {
		codec Codec
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "good", args: args{codec{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterCodec(tt.args.codec)

			if got := GetCodec(tt.args.codec.Name()); !reflect.DeepEqual(got, tt.args.codec) {
				t.Errorf("GetCodec() = %v, want %v", got, tt.args.codec)
			}
		})
	}
}
