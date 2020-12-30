package retry

import (
	"context"
	"time"
)

var (
	defaultMaxAttempts  uint    = 10
	defaultFactor       float64 = 2
	defaultIntensity            = time.Millisecond * 500
	defaultMaxRetryTime         = time.Second * 10
)

// Interface represents Retry options
type Option interface {
	apply(*retryOptions)
}

type funcOptions struct {
	f func(opt *retryOptions)
}

func (fo *funcOptions) apply(o *retryOptions) {
	fo.f(o)
}

func newFuncOptions(f func(opt *retryOptions)) *funcOptions {
	return &funcOptions{
		f: f,
	}
}

type retryOptions struct {
	backoff *backoff
	ctx     context.Context
}

func defaultOptions() retryOptions {
	return retryOptions{
		backoff: NewBackoff(defaultMaxAttempts, defaultFactor, defaultMaxRetryTime, defaultIntensity, false),
		ctx:     context.Background(),
	}
}

// Configure max attempts
func WithMaxAttempts(num int) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.maxAttempts = uint(num)
	})
}


// Configure first attempt
func WithAttempts(num int) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.attemptNum = uint(num)
	})
}

// Configure factor for exponential backoff
func WithFactor(num float64) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.factor = num
	})
}

// Configure max retry time
func WithMaxRetryTime(t time.Duration) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.maxRetryTime = t
	})
}

// Configure retry intensity
func WithRetryIntensity(t time.Duration) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.intensity = t
	})
}

// Pass context
func WithContext(ctx context.Context) Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.ctx = ctx
	})
}

// Use linear backoff
func WithLinearFunc() Option {
	return newFuncOptions(func(opt *retryOptions) {
		opt.backoff.isLinear = true
	})
}
