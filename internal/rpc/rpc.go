package rpc

import (
	"fmt"
)

type RpcConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:"8085"`
}

func (config RpcConfig) GetAddress() string {
	return fmt.Sprintf("%s%s", config.Host, config.Port)
}
