package mobilerpc

import (
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// TODO
// Что то с аккаунтом (откуда брать, через что сохранять)
// Тесты
// Придумать интерфейс
// Подумать стоит ли оставлять функции makeConnection

type mobileRpc struct {
	config      Config
	grpcContext RpcContext
}

// Get charts from external api
func (rpc *mobileRpc) Charts(ctx context.Context, chart Category) ([]string, error) {
	conn, err := rpc.MakeConnection()
	if err != nil {
		return nil, err
	}
	client := charts.NewGuardClient(conn)

	cat, subcat := chart.Split()
	resp, err := client.TopCharts(ctx, &charts.ChartsRequest{
		Cat:     cat,
		SubCat:  subcat,
		Account: rpc.config.RpcAccount,
	})
	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, fmt.Errorf("rpc call response with %d status code", resp.ErrCode)
	}
	rpc.checkTokens(resp.Account)

	return resp.Ids, nil
}

// Make new grpc connection to external api
func (rpc *mobileRpc) MakeConnection() (*grpc.ClientConn, error) {
	if rpc.grpcContext == nil {
		rpc.grpcContext = NewGrpcContext(
			fmt.Sprintf("%s:%s", rpc.config.Address, rpc.config.Port),
			15,
		)
	}

	conn, err := rpc.grpcContext.WakeUp()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Check token and gsfid and set them to local config struct
// if they are empty
func (rpc *mobileRpc) checkTokens(acc *charts.Account) {
	if rpc.config.RpcAccount.GsfId == 0 || rpc.config.RpcAccount.Token == "" {
		rpc.config.RpcAccount.Token = acc.Token
		rpc.config.RpcAccount.GsfId = acc.GsfId
	}
}

// Create new instance of mobile rpc api
func New(config Config) *mobileRpc {
	return &mobileRpc{
		config: config,
	}
}
