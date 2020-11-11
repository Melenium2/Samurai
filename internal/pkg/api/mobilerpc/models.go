package mobilerpc

import (
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"fmt"
	"strings"
)

type Category string

func (c Category) Split() (string, string) {
	splited := strings.Split(string(c), "|")
	if len(splited) > 2 {
		panic("invalid category")
	}
	return splited[0], splited[1]
}

func NewCategory(cat, subcat string) Category {
	return Category(strings.ToLower(fmt.Sprintf("%s|%s", cat, subcat)))
}

type Account struct {
	Login    string
	Password string
	GsfId    int
	Token    string
	Locale   string
	Proxy    *Proxy
	Device   string
}

func (a Account) ForGrpc() *charts.Account {
	return &charts.Account{
		Login:    a.Login,
		Password: a.Password,
		GsfId:    int64(a.GsfId),
		Token:    a.Token,
		Locale:   a.Locale,
		Proxy:    a.Proxy.ForGrpc(),
		Device:   a.Device,
	}
}

func (a *Account) Fill(account *charts.Account) {
	a.Login = account.Login
	a.Password = account.Password
	a.GsfId = int(account.GsfId)
	a.Token = account.Token
	a.Login = account.Locale
	p := &Proxy{}
	p.Fill(account.Proxy)
	a.Proxy = p
	a.Device = account.Device
}

type Proxy struct {
	Http  string
	Https string
}

func (p Proxy) ForGrpc() *charts.Proxy {
	return &charts.Proxy{
		Http: p.Http,
		Https: p.Https,
	}
}

func (p *Proxy) Fill(proxy *charts.Proxy) {
	p.Http = proxy.Http
	p.Https = proxy.Https
}
