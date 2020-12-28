package inhuman

import (
	"Samurai/config"
	"strings"
)

type Config struct {
	Url        string
	Key        string
	Hl         string
	Gl         string
	ItemsCount int
}

func FromConfig(config config.Config) Config {
	splitedLang := strings.Split(config.App.Lang, "_")
	hl, gl := strings.ToLower(splitedLang[0]), strings.ToLower(splitedLang[1])
	return Config{
		Url:        config.Api.Url,
		Key:        config.Api.Key,
		Hl:         hl,
		Gl:         gl,
		ItemsCount: config.App.ItemsCount,
	}
}
