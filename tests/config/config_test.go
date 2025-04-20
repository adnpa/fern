package config

import (
	"fmt"
	"testing"

	"github.com/adnpa/fern/pkg/common/config"
	"github.com/adnpa/fern/pkg/common/config/file"
)

type Conf struct {
	Abc string `json:"abc,omitempty"`
	Hij string `json:"hij,omitempty"`
}

func Test_config_Load(t *testing.T) {
	tests := []struct {
		name    string
		c       config.Config
		wantErr bool
	}{
		{name: "json file", c: config.NewConfig(config.WithSource(file.NewSource("../../test_data/config/test.json")))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Load(); (err != nil) != tt.wantErr {
				t.Errorf("config.Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			val := tt.c.Value("abc")
			t.Log(val)

			var conf Conf
			err := tt.c.Scan(&conf)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(conf)
		})
	}
}
