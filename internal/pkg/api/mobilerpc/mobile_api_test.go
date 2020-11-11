package mobilerpc_test

import (
	"Samurai/internal/pkg/api/mobilerpc"
	charts "Samurai/internal/pkg/api/mobilerpc/proto"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"log"
	"testing"
)

type mobile_mock struct {
	mobilerpc.Config

	ResponseCode int32
}

func (m *mobile_mock) Charts(ctx context.Context, chart mobilerpc.Category) ([]string, error) {
	cat, subcat := chart.Split()

	log.Print(cat, "",subcat)

	if m.ResponseCode > 0 {
		if m.ResponseCode == 1001 {
			m.RpcAccount.GsfId = 0
			m.RpcAccount.Token = ""
		}
		return nil, fmt.Errorf("rpc call response with %d status code", m.ResponseCode)
	}

	return []string{"1", "2", "3"}, nil
}

func (m mobile_mock) MakeConnection() (*grpc.ClientConn, error) {
	return grpc.Dial("")
}

func TestMobileRpc_MakeConnection_ShouldReturnNewConnection_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{})

	conn, err := api.MakeConnection()
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestMobileRpc_Charts_NoError(t *testing.T) {
	var tests = []struct {
		name            string
		responseCode    int32
		expectedObject  []string
		expectedAccount mobilerpc.Account
		err             bool
	}{
		{
			name:           "Should response random code error from grpc",
			responseCode:   10000,
			expectedObject: []string{},
		},
		{
			name:            "should response 1001 error from grpc and clear account settings ",
			responseCode:    1001,
			expectedObject:  []string{},
			expectedAccount: mobilerpc.Account{Token: "", GsfId: 0},
		},
		{
			name:            "should return bundles",
			responseCode:    0,
			expectedObject:  []string{"1", "2", "3"},
			expectedAccount: mobilerpc.Account{GsfId: 100, Token: "123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := mobile_mock{
				Config:       mobilerpc.Config{RpcAccount: &mobilerpc.Account{GsfId: 100, Token: "123"}},
				ResponseCode: tt.responseCode,
			}
			resp, err := api.Charts(context.Background(), mobilerpc.NewCategory("FINANCE", "apps_topselling_free"))
			if tt.responseCode > 0 {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedObject, []string{})
				assert.Equal(t, &tt.expectedAccount, &charts.Account{Token: "", GsfId: 0})
				return
			}
			assert.Equal(t, tt.expectedObject, resp)
			assert.Equal(t, &tt.expectedAccount, &charts.Account{Token: "123", GsfId: 100})
		})
	}
}

func TestMobileRpc_MakeConnection_ShouldEstablishConnection_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{
		Address: "localhost",
		Port: "1000",
		RpcAccount: &mobilerpc.Account{
			Login:    "",
			Password: "",
			GsfId:    4024815430645318922,
			Token:    "2QfEo-v0LuEV1bs3Ei-QskzGiGG5cYRf69NGlq8DYtpOd7nkbXLPrgayovdKhYsz91baVw.",
			Locale:   "ru_RU",
			Proxy:    nil,
			Device:   "whyred",
		},
	})
	conn, err := api.MakeConnection()
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestMobileRpc_Charts_ShouldMakeRequestAndGetResponse_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{
		Address: "localhost",
		Port: "1000",
		RpcAccount: &mobilerpc.Account{
			Login:    "",
			Password: "",
			GsfId:    4024815430645318922,
			Token:    "2QfEo-v0LuEV1bs3Ei-QskzGiGG5cYRf69NGlq8DYtpOd7nkbXLPrgayovdKhYsz91baVw.",
			Locale:   "ru_RU",
			Proxy:    nil,
			Device:   "whyred",
		},
	})
	resp, err := api.Charts(context.Background(), mobilerpc.NewCategory("GAME_ACTION", "apps_topselling_free"))
	assert.NoError(t, err)
	assert.Greater(t, len(resp), 0)
}

func TestMobileRpc_Charts_ShouldGetEmptyResponse_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{
		Address: "localhost",
		Port: "1000",
		RpcAccount: &mobilerpc.Account{
			Login:    "",
			Password: "",
			GsfId:    4024815430645318922,
			Token:    "2QfEo-v0LuEV1bs3Ei-QskzGiGG5cYRf69NGlq8DYtpOd7nkbXLPrgayovdKhYsz91baVw.",
			Locale:   "ru_RU",
			Proxy:    nil,
			Device:   "whyred",
		},
	})
	resp, err := api.Charts(context.Background(), mobilerpc.NewCategory("", "ap1ps_topselling_free"))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp))
}

func TestMobileRpc_Charts_ShouldGetErrorCode1001_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{
		Address: "localhost",
		Port: "1000",
		RpcAccount: &mobilerpc.Account{
			Login:    "",
			Password: "",
			GsfId:    0,
			Token:    "2QfEo-v0LuEV1bs3Ei-QskzGiGG5cYRf69NGlq8DYtpOd7nkbXLPrgayovdKhYsz91baVw.",
			Locale:   "ru_RU",
			Proxy:    nil,
			Device:   "whyred",
		},
	})
	resp, err := api.Charts(context.Background(), mobilerpc.NewCategory("GAME_ACTION", "apps_topselling_free"))
	assert.Error(t, err)
	assert.Equal(t, 0, len(resp))
	assert.Equal(t, int64(0), api.Config.RpcAccount.GsfId)
	assert.Equal(t, "", api.Config.RpcAccount.Token)
}

func TestMobileRpc_Charts_ShouldLoginAndAddNewTokenAndId_NoError(t *testing.T) {
	api := mobilerpc.New(mobilerpc.Config{
		Address: "localhost",
		Port: "1000",
		RpcAccount: &mobilerpc.Account{
			Login:    "ceciliamcalistervugt93@gmail.com",
			Password: "Hbibcxzauig",
			GsfId:    0,
			Token:    "",
			Locale:   "ru_RU",
			Proxy:    &mobilerpc.Proxy{
				Http:  "http://STqthJ:2odx6V@45.132.21.233:8000",
				Https: "https://STqthJ:2odx6V@45.132.21.233:8000",
			},
			Device:   "whyred",
		},
	})
	resp, err := api.Charts(context.Background(), mobilerpc.NewCategory("FINANCE", "apps_topselling_free"))
	assert.NoError(t, err)
	assert.Greater(t, len(resp), 0)
	assert.NotEqual(t, int64(0), api.Config.RpcAccount.GsfId)
	assert.NotEqual(t, "", api.Config.RpcAccount.Token)
}