package config

import (
	"Samurai/internal/pkg/api/mobilerpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

//Database config
type DBConfig struct {
	Database string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Schema   string `yaml:"schema"`
}

// Api config
type ApiConfig struct {
	Url         string `yaml:"url"`
	Key         string `yaml:"key"`
	GrpcAddress string `yaml:"grpc_address"`
	GrpcPort    string `yaml:"grpc_port"`
	GrpcAccount *mobilerpc.Account
}

// Main config
type AppConfig struct {
	Bundle    string        `yaml:"bundle"`
	Period    int           `yaml:"period"`
	Intensity time.Duration `yaml:"intensity"`
	Lang      string        `yaml:"lang"`
	Keywords  []string      `yaml:",flow"`
}

//Application config
type Config struct {
	Api      ApiConfig `yaml:"api"`
	Database DBConfig  `yaml:"database"`
	App      AppConfig `yaml:"app"`

	Envs []string `yaml:",flow"`
}

/**
Load system environment variables from given array
*/
func loadEnvs(e ...string) map[string]string {
	envs := make(map[string]string)

	for _, k := range e {
		envs[k] = os.Getenv(k)
	}

	return envs
}

//Create new instance of app config with given path
func New(p ...string) Config {
	path := "./dev.yml"
	if len(p) > 0 {
		path = p[0]
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := Config{
		Api: ApiConfig{
			GrpcAccount: &mobilerpc.Account{},
		},
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	envs := loadEnvs(config.Envs...)

	v, ok := envs["api_key"]
	if ok && v != "" {
		config.Api.Key = v
	}
	v, ok = envs["db_pass"]
	if ok && v != "" {
		config.Database.Password = v
	}
	v, ok = envs["db_user"]
	if ok && v != "" {
		config.Database.User = v
	}

	return config
}
