package mobilerpc

import charts "Samurai/internal/pkg/api/mobilerpc/proto"

type Config struct {
	Address    string
	Port       string
	RpcAccount *charts.Account
}
