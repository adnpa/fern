// source from file,implement source interface
package file

import (
	"os"
	"testing"

	"github.com/adnpa/fern/pkg/common/config"
)

func Test_file_Load(t *testing.T) {
	t.Log(os.Getwd())

	type fields struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*config.KV
		wantErr bool
	}{
		{name: "json", fields: fields{path: "../../../../test_data/config/test.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &file{
				path: tt.fields.path,
			}
			got, err := f.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("file.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, v := range got {
				t.Log(v.Key, string(v.Value), v.Format)
			}
			t.Log(err)
		})
	}
}
