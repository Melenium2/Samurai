package mobilerpc_test

import (
	"Samurai/internal/pkg/api/mobilerpc"
	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestGrpcContext_Update_ShouldStoreNewValueToTimer_NoError(t *testing.T) {
	var tests = []struct {
		name string
		in   int
		out  int
	}{
		{
			name: "If value equal zero up to 30",
			in:   0,
			out:  30,
		},
		{
			name: "If value less then zero up to 30",
			in:   -30,
			out:  30,
		},
		{
			name: "If value more then 30 round to 30",
			in:   60,
			out:  30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mobilerpc.GrpcContext{
				Period: 30,
			}
			if tt.in < 100 {
				ctx.Timeout = atomic.NewInt32(int32(tt.in))
			}
			ctx.Update()
			assert.Equal(t, int32(tt.out), ctx.Timeout.Load())
		})
	}
}

func TestGrpcContext_Destroy_ShouldClearTimeoutAndSetNilToConnection_NoError(t *testing.T) {
	grpcConn, err := grpc.Dial("", grpc.WithInsecure())
	assert.NoError(t, err)
	var tests = []struct {
		name         string
		time         int32
		expectedTime int32
		conn         *grpc.ClientConn
	}{
		{
			name:         "Expected error coz conn already nil",
			time:         30,
			expectedTime: 0,
			conn:         nil,
		},
		{
			name:         "Expected time = 0 and connection is nil",
			time:         30,
			expectedTime: 0,
			conn:         grpcConn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mobilerpc.GrpcContext{}
			ctx.Timeout = atomic.NewInt32(tt.time)
			ctx.Conn = tt.conn

			ctx.Destroy()

			assert.Nil(t, ctx.Conn)
			assert.Equal(t, int32(0), ctx.Timeout.Load())
		})
	}
}

func TestGrpcContext_WakeUp_ShouldPassAllTableTests_NoError(t *testing.T) {
	grpcConn, _ := grpc.Dial("", grpc.WithInsecure())

	var tests = []struct {
		name         string
		conn         *grpc.ClientConn
		timeout      bool
		expectedTime int32
	} {
		{
			name: "Context already has connection. Return conn.",
			timeout: false,
			conn: grpcConn,
		},
		{
			name: "Create connection. Connection not nil.",
			timeout: false,
			expectedTime: 4,
		},
		{
			name: "Create connection. Wait until time os over.",
			timeout: true,
			expectedTime: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mobilerpc.GrpcContext{
				Timeout: atomic.NewInt32(0),
				Period: 5,
			}
			if tt.conn != nil {
				ctx.Conn = tt.conn
			}
			conn, err := ctx.WakeUp()
			assert.NoError(t, err)
			assert.NotNil(t, conn)
			assert.NotNil(t, ctx.Conn)
			if tt.conn == nil {
				assert.Greater(t, ctx.Timeout.Load(), tt.expectedTime)
			}


			if tt.timeout {
				time.Sleep(time.Second * 6)
				assert.Nil(t, ctx.Conn)
				assert.Equal(t, int32(0), ctx.Timeout.Load())
			}
		})
	}
}

func TestGrpcContext_New_ShouldCreateNewContextCheckConnectionAndDestroyHim_NoError(t *testing.T) {
	ctx := mobilerpc.NewGrpcContext("", 15)
	conn, err := ctx.WakeUp()
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NotNil(t, ctx.Conn)

	time.Sleep(time.Second * 3)

	ctx.Destroy()

	assert.Nil(t, ctx.Conn)
	assert.LessOrEqual(t, int32(0), ctx.Timeout.Load())
}
