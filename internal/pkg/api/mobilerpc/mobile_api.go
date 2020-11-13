package mobilerpc

import (
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// TODO
// Что то с аккаунтом (откуда брать, через что сохранять)
// Тесты

type ChartApi interface {
	Charts(ctx context.Context, chart Category) ([]string, error)
}

type mobileRpc struct {
	Config

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
		Account: rpc.RpcAccount.ForGrpc(),
	})
	if err != nil {
		log.Print("Where top charts")
		return nil, err
	}

	rpc.checkTokens(resp.Account)
	if resp.ErrCode > 0 {
		if resp.ErrCode == 1001 {
			rpc.RpcAccount.GsfId = 0
			rpc.RpcAccount.Token = ""
		}
		return nil, fmt.Errorf("rpc call response with %d status code", resp.ErrCode)
	}

	return resp.Ids, nil
}

// Make new grpc connection to external api
func (rpc *mobileRpc) MakeConnection() (*grpc.ClientConn, error) {
	if rpc.grpcContext == nil {
		rpc.grpcContext = NewGrpcContext(
			fmt.Sprintf("%s:%s", rpc.Address, rpc.Port),
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
	if acc != nil && (rpc.RpcAccount.GsfId == 0 || rpc.RpcAccount.Token == "") {
		rpc.RpcAccount.Token = acc.Token
		rpc.RpcAccount.GsfId = int(acc.GsfId)
	}
}

// Create new instance of mobile rpc api
func New(config Config) *mobileRpc {
	return &mobileRpc{
		Config: config,
	}
}
