package config

import (
	"reflect"
	"testing"
	"time"
)

func Test_atomicValue_typeAssertError(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.typeAssertError(); (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.typeAssertError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_atomicValue_Bool(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("atomicValue.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Int(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("atomicValue.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Slice(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    []Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atomicValue.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Map(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    map[string]Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Map()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atomicValue.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Float(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Float()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Float() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("atomicValue.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_String(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("atomicValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Duration(t *testing.T) {
	tests := []struct {
		name    string
		v       *atomicValue
		want    time.Duration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Duration()
			if (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Duration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atomicValue.Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicValue_Scan(t *testing.T) {
	type args struct {
		obj any
	}
	tests := []struct {
		name    string
		v       *atomicValue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Scan(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("atomicValue.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_errValue_Bool(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("errValue.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Int(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("errValue.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Float(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Float()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Float() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("errValue.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Duration(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    time.Duration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Duration()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Duration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errValue.Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_String(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("errValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Scan(t *testing.T) {
	type args struct {
		in0 any
	}
	tests := []struct {
		name    string
		v       errValue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Scan(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("errValue.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_errValue_Load(t *testing.T) {
	tests := []struct {
		name string
		v    errValue
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Load(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errValue.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Store(t *testing.T) {
	type args struct {
		in0 any
	}
	tests := []struct {
		name string
		v    errValue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Store(tt.args.in0)
		})
	}
}

func Test_errValue_Slice(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    []Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errValue.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errValue_Map(t *testing.T) {
	tests := []struct {
		name    string
		v       errValue
		want    map[string]Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Map()
			if (err != nil) != tt.wantErr {
				t.Errorf("errValue.Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errValue.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
