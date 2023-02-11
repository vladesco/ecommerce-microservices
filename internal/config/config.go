package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
	"github.com/vladesco/ecommerce-microservices/internal/database"
	"github.com/vladesco/ecommerce-microservices/internal/logger"
	"github.com/vladesco/ecommerce-microservices/internal/rpc"
	"github.com/vladesco/ecommerce-microservices/internal/web"
)

type Environment string

const (
	Production    Environment = "prod"
	PreProduction Environment = "preprod"
	Develop       Environment = "dev"
)

type AppConfig struct {
	Environment
	LogLevel        logger.LogLevel
	Database        database.DatabaseConfig
	Rpc             rpc.RpcConfig
	Web             web.WebConfig
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"15s"`
}

func GetAppConfig(environment Environment) (config AppConfig, err error) {
	if err = dotenv.Load(
		dotenv.EnvironmentFiles(string(environment)),
	); err != nil {
		return
	}

	err = envconfig.Process("", &config)

	return
}
