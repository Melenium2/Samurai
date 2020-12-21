package mobilerpc

import (
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"strings"
)

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

type Proxy struct {
	Http  string
	Https string
}

func NewProxy(proxy string) *Proxy {
	p := strings.Split(proxy, "//")

	return &Proxy{
		Http:  "http://" + p[1],
		Https: "https://" + p[1],
	}
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
