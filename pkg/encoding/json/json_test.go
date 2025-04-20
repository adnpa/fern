package json

import (
	"reflect"
	"testing"
)

type TestObj struct {
	Field1 string
	Field2 string
	Field3 string
}

func Test_codec_RoundTrip(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		c    codec
		args args
		// want    []byte
		wantErr bool
	}{
		{name: "good", args: args{TestObj{Field1: "f1", Field2: "f2"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := codec{}
			got, err := c.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("codec.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var result TestObj
			if err = c.Unmarshal(got, &result); (err != nil) != tt.wantErr {
				t.Errorf("codec.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.args.v, result) {
				t.Errorf("Round trip mismatch:\noriginal: %+v\nresult: %+v", tt.args.v, result)
			}
		})
	}
}

func Test_codec_Name(t *testing.T) {
	tests := []struct {
		name string
		c    codec
		want string
	}{
		{name: "good", want: "json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := codec{}
			if got := c.Name(); got != tt.want {
				t.Errorf("codec.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
