package config

import (
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"Samurai/internal/pkg/api/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type TrackingType uint

// Type of stores tracker can track
const (
	AppStore TrackingType = iota
	GooglePlay
)

// Account represent user account in external mobile api
type Account struct {
	Login    string
	Password string
	GsfId    int
	Token    string
	Locale   string
	Proxy    *Proxy
	Device   string
}

// Convert Account to *charts.Account
func (a Account) ForGrpc() *charts.Account {
	var p *charts.Proxy
	if a.Proxy != nil {
		p = a.Proxy.ForGrpc()
	}
	return &charts.Account{
		Login:    a.Login,
		Password: a.Password,
		GsfId:    int64(a.GsfId),
		Token:    a.Token,
		Locale:   a.Locale,
		Proxy:    p,
		Device:   a.Device,
	}
}

// Fill fields from *charts.Account
func (a *Account) Fill(account *charts.Account) {
	a.Login = account.Login
	a.Password = account.Password
	a.GsfId = int(account.GsfId)
	a.Token = account.Token
	a.Login = account.Locale
	p := &Proxy{}
	if account.Proxy != nil {
		p.Fill(account.Proxy)
	}
	a.Proxy = p
	a.Device = account.Device
}

// Proxy represents proxy object
type Proxy struct {
	Http  string
	Https string
}

// Construct new Proxy struct by proxy string
// proxy string must be like http(or https)://*****:*****@ipaddress:port
func NewProxy(proxy string) *Proxy {
	p := strings.Split(proxy, "//")

	return &Proxy{
		Http:  "http://" + p[1],
		Https: "https://" + p[1],
	}
}

// Convert Proxy to *charts.Proxy
func (p Proxy) ForGrpc() *charts.Proxy {
	return &charts.Proxy{
		Http:  p.Http,
		Https: p.Https,
	}
}

// Fill *Proxy from *charts.Proxy
func (p *Proxy) Fill(proxy *charts.Proxy) {
	p.Http = proxy.Http
	p.Https = proxy.Https
}

//Database config
type DBConfig struct {
	Database string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Schema   string `yaml:"schema"`
}

// ApiConfig
type ApiConfig struct {
	Url         string `yaml:"url"`
	Key         string `yaml:"key"`
	GrpcAddress string `yaml:"grpc_address"`
	GrpcPort    string `yaml:"grpc_port"`
	GrpcAccount Account
}

// Main config AppConfig
type AppConfig struct {
	Bundle      string        `yaml:"bundle"`
	Period      int           `yaml:"period"`
	Intensity   time.Duration `yaml:"intensity"`
	Lang        string        `yaml:"lang"`
	Keywords    []string      `yaml:",flow"`
	ExternalLog string        `yaml:"external_logger"`
	// Keywords or application request count from external api
	ItemsCount int `yaml:"count"`
	Categories models.Collection
	OnlyMeta   bool

}

// Config struct of application config
type Config struct {
	Api      ApiConfig `yaml:"api"`
	Database DBConfig  `yaml:"database"`
	App      AppConfig `yaml:"app"`

	Envs []string `yaml:",flow"`
}

// View print full config
func (c Config) View() {
	log.Print("____START")
	log.Print("Environments: ", c.Envs)

	log.Print("---------------------------")

	log.Print("***DEPENDENCY***")
	log.Print("\tApiUrl: ", c.Api.Url)
	log.Print("\tGrpcPort: ", c.Api.GrpcPort)
	log.Print("\tGrpcAddress: ", c.Api.GrpcAddress)

	log.Print("---------------------------")

	log.Print("***ACCOUNT***")
	log.Print("\tLogin: ", c.Api.GrpcAccount.Login)
	log.Print("\tPassword: ", c.Api.GrpcAccount.Password)
	log.Print("\tToken: ", c.Api.GrpcAccount.Token)
	log.Print("\tGSFID: ", c.Api.GrpcAccount.GsfId)
	log.Print("\tDevice: ", c.Api.GrpcAccount.Device)
	log.Print("\tProxy: ", c.Api.GrpcAccount.Proxy)
	log.Print("\tLocale: ", c.Api.GrpcAccount.Locale)

	log.Print("---------------------------")

	log.Print("***APPLICATION***")
	log.Print("\tLanguage: ", c.App.Lang)
	log.Print("\tIntensity: ", c.App.Intensity)
	log.Print("\tPeriod: ", c.App.Period)
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

// ladEnvs load system envs from given array of keys and returns
// map of sysenv key : value
func loadEnvs(e ...string) map[string]string {
	envs := make(map[string]string)

	for _, k := range e {
		envs[k] = os.Getenv(k)
	}

	return envs
}

// New Create new instance of app config with given path to (..config..).yml
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
		App: AppConfig{
			Categories: models.CategoriesGoogle,
		},
		Api: ApiConfig{
			GrpcAccount: Account{},
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
