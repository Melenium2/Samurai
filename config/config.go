package config

import (
	"Samurai/internal/pkg/api/mobilerpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
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

// Print application full config
func (c Config) View() {
	log.Print("____START")
	log.Print("Environments: ", c.Envs)

	log.Print("---------------------------")

	log.Print("***ACCOUNT***")
	log.Print("\tLogin:", c.Api.GrpcAccount.Login)
	log.Print("\tPassword: ", c.Api.GrpcAccount.Password)
	log.Print("\tToken: ", c.Api.GrpcAccount.Token)
	log.Print("\tGSFID: ", c.Api.GrpcAccount.GsfId)
	log.Print("\tDevice: ", c.Api.GrpcAccount.Device)
	log.Print("\tProxy: ", c.Api.GrpcAccount.Proxy)
	log.Print("\tLocale: ", c.Api.GrpcAccount.Locale)

	log.Print("---------------------------")

	log.Print("***APPLICATION***")
	log.Print("\tLanguage: ", c.App.Lang)
	log.Print("\tIntensity ", c.App.Intensity)
	log.Print("\tPeriod ", c.App.Period)
	log.Print("\tKeywords: ", c.App.Keywords)
	log.Print("\tBundle: ", c.App.Bundle)

	log.Print("---------------------------")

	log.Print("***DATABASE***")
	log.Print("\tAddress: ", c.Database.Address)
	log.Print("\tPort: ", c.Database.Port)
	log.Print("\tUser: ", c.Database.User)
	log.Print("\tPassword: ", c.Database.Password)
	log.Print("\tDatabase: ", c.Database.Database)
	log.Print("\tSchema file: ", c.Database.Schema)
	log.Print("____END")
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
	v, ok = envs["grpc_address"]
	if ok && v != "" {
		config.Api.GrpcAddress = v
	}
	v, ok = envs["grpc_port"]
	if ok && v != "" {
		config.Api.GrpcPort = v
	}


	return config
}
