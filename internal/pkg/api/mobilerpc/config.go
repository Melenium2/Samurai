package mobilerpc

import "Samurai/config"

// Config to grpc connection
type Config struct {
	Address    string
	Port       string
	RpcAccount config.Account
}

// Generate Config from config.Config
func FromConfig(config config.Config) Config {
	return Config{
		config.Api.GrpcAddress, config.Api.GrpcPort, config.Api.GrpcAccount,
	}
}
