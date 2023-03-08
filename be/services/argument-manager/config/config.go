package config

import (
	"os"
)

const (
	// Environments
	ProdEnv = "prod"
	TestEnv = "test"

	// Env vars
	appEnv      = "APP_ENV" // intentionally without prefix
	GrpcPortEnv = "ARGUMENT_ACCESS_GRPC_PORT"
)

type Config struct {
	Env      string
	GrpcPort string
}

func NewConfig() Config {
	c := Config{}

	// Set defaults
	c.Env = TestEnv
	c.GrpcPort = "9002"

	return c
}

func GetConfig() *Config {
	c := NewConfig()

	env, present := os.LookupEnv(appEnv)
	if present {
		c.Env = env
	}

	GrpcPort, present := os.LookupEnv(GrpcPortEnv)
	if present {
		c.GrpcPort = GrpcPort
	}

	return &c
}
