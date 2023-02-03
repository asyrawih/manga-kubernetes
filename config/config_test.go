package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "should oke load config file",
			want: &Config{
				Mysql: Mysql{
					Mysqldatabase: "localhost",
					Mysqlhostname: "root",
					Mysqlpassword: "root",
					Mysqlusername: "manga",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadConfig("config.json")
			assert.NotNil(t, got)
		})
	}
}
