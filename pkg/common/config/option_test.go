package config

import (
	"reflect"
	"testing"
)

func TestNewOptions(t *testing.T) {
	tests := []struct {
		name string
		want *Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSource(t *testing.T) {
	type args struct {
		s []Source
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithSource(tt.args.s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDecoder(t *testing.T) {
	type args struct {
		d Decoder
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDecoder(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMergeFunc(t *testing.T) {
	type args struct {
		m Merge
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMergeFunc(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMergeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultDecoder(t *testing.T) {
	type args struct {
		cfg    *KV
		target map[string]any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultDecoder(tt.args.cfg, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("defaultDecoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultMerge(t *testing.T) {
	type args struct {
		dst any
		src any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultMerge(tt.args.dst, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("defaultMerge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
