package web

import (
	"fmt"
)

type WebConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:"8080"`
}

func (config WebConfig) GetAddress() string {
	return fmt.Sprintf("%s%s", config.Host, config.Port)
}
