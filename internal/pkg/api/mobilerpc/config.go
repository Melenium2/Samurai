package mobilerpc

import "Samurai/config"

type Config struct {
	Address    string
	Port       string
	RpcAccount config.Account
}

func FromConfig(config config.Config) Config {
	return Config{
		config.Api.GrpcAddress, config.Api.GrpcPort, config.Api.GrpcAccount,
	}
}
