package config

import (
	"reflect"
	"testing"

	"github.com/adnpa/fern/pkg/common/config"
)

func TestNewReader(t *testing.T) {
	type args struct {
		opts config.Options
	}
	tests := []struct {
		name string
		args args
		want config.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.NewReader(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_reader_Merge(t *testing.T) {
// 	type args struct {
// 		kvs []*KV
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *reader
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.r.Merge(tt.args.kvs...); (err != nil) != tt.wantErr {
// 				t.Errorf("reader.Merge() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_reader_Value(t *testing.T) {
// 	type args struct {
// 		k string
// 	}
// 	tests := []struct {
// 		name  string
// 		r     *reader
// 		args  args
// 		want  Value
// 		want1 bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := tt.r.Value(tt.args.k)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("reader.Value() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("reader.Value() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// func Test_reader_Source(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		r       *reader
// 		want    []byte
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.r.Source()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("reader.Source() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("reader.Source() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_reader_Resolve(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		r       *reader
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.r.Resolve(); (err != nil) != tt.wantErr {
// 				t.Errorf("reader.Resolve() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
