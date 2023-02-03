package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

type Mysql struct {
	Mysqldatabase string `json:"mysqldatabase,omitempty"`
	Mysqlhostname string `json:"mysqlhostname,omitempty"`
	Mysqlpassword string `json:"mysqlpassword,omitempty"`
	Mysqlusername string `json:"mysqlusername,omitempty"`
}

type Config struct {
	Mysql `json:"mysql,omitempty"`
}

func LoadConfig(path string) *Config {

	var config *Config
	b, err := os.ReadFile(path)

	if err != nil {
		log.Err(err).Caller().Msg("")
	}

	if err := json.Unmarshal(b, &config); err != nil {
		log.Err(err).Caller().Msg("")
	}

	return config
}
