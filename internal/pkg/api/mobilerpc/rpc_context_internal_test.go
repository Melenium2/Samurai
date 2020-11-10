package mobilerpc

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestGrpcContext_tick_ShouldCallDestroyAndClearConnectionAfter3Sec_NoError(t *testing.T) {
	grpcConn, err := grpc.Dial("", grpc.WithInsecure())
	assert.NoError(t, err)
	ctx := GrpcContext{
		Timeout: atomic.NewInt32(3),
		Conn: grpcConn,
	}

	ctx.tick()

	assert.Nil(t, ctx.Conn)
	assert.Equal(t, int32(0), ctx.Timeout.Load())
}

func TestGrpcContext_tick_ShouldCallDestroyAsyncAfter5Sec_NoError(t *testing.T) {
	grpcConn, err := grpc.Dial("", grpc.WithInsecure())
	assert.NoError(t, err)
	ctx := GrpcContext{
		Timeout: atomic.NewInt32(3),
		Conn: grpcConn,
	}

	go ctx.tick()

	time.Sleep(time.Second * 4)

	assert.Nil(t, ctx.Conn)
	assert.Equal(t, int32(0), ctx.Timeout.Load())
}
