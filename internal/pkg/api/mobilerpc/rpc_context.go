package mobilerpc

import (
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"log"
	"time"
)

// RpcContext save grpc connection for period of time
// then clear connection. We can wake up connection
// if needed
type RpcContext interface {
	Update()
	Destroy()
	WakeUp() (*grpc.ClientConn, error)
}

// GrpcContext struct implementation of RpcContext
type GrpcContext struct {
	Conn    *grpc.ClientConn
	Timeout *atomic.Int32
	Address string
	period  int32
}

// Update timeout if user renews wanna connect longer
func (gctx *GrpcContext) Update() {
	gctx.Timeout.Store(gctx.period)
}

// Destroy connection
func (gctx *GrpcContext) Destroy() {
	if gctx.Conn != nil {
		err := gctx.Conn.Close()
		if err != nil {
			log.Print(err)
		}
	}

	gctx.Conn = nil
	gctx.Timeout.Store(0)
}

// WakeUp connection or just return connection
func (gctx *GrpcContext) WakeUp() (*grpc.ClientConn, error) {
	if gctx.Conn != nil {
		return gctx.Conn, nil
	}
	conn, err := grpc.Dial(gctx.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	gctx.Conn = conn
	gctx.Update()
	go gctx.tick()

	return conn, nil
}

// Timer until connection closed
func (gctx *GrpcContext) tick() {
	for {
		t := gctx.Timeout.Dec()
		if t < 0 {
			break
		}
		time.Sleep(time.Second)
	}

	gctx.Destroy()
}

// Create new instance of context and automatically create new connection
// if flag not nil
func NewGrpcContext(address string, period int, connect ...bool) *GrpcContext {
	c := &GrpcContext{
		Address: address,
		Timeout: atomic.NewInt32(int32(period)),
	}

	if len(connect) > 0 {
		_, err := c.WakeUp()
		if err != nil {
		panic(err)
	}
	}

	return c
}
