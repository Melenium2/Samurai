package retry

import "time"

var (
	defaultMaxAttempts uint = 10
	defaultAttemptNum  uint = 0
	defaultFactor      uint = 2
	defaultMin              = time.Millisecond * 500
	defaultMax              = time.Second * 10
)

type Options interface {
	apply(*options)
}

type funcOptions struct {
	f func(opt *options)
}

func (fo *funcOptions) apply(o *options) {
	fo.f(o)
}

func newFuncOptions(f func(opt *options)) *funcOptions {
	return &funcOptions{
		f: f,
	}
}

type options struct {
	maxAttempts, attemptNum, factor uint
	min, max                        time.Duration
}

func DefaultOptions() Options {
	return newFuncOptions(func(opt *options) {
		opt.attemptNum = defaultAttemptNum
		opt.factor = defaultFactor
		opt.maxAttempts = defaultMaxAttempts
		opt.min = defaultMin
		opt.max = defaultMax
	})
}

func WithMaxAttempts(num int) Options {
	return newFuncOptions(func(opt *options) {
		opt.maxAttempts = uint(num)
	})
}

func WithAttempts(num int) Options {
	return newFuncOptions(func(opt *options) {
		opt.attemptNum = uint(num)
	})
}

func WithFactor(num int) Options {
	return newFuncOptions(func(opt *options) {
		opt.factor = uint(num)
	})
}

func WithMaxRetryTime(t time.Duration) Options {
	return newFuncOptions(func(opt *options) {
		opt.max = t
	})
}

func WithRetryIntensivity(t time.Duration) Options {
	return newFuncOptions(func(opt *options) {
		opt.min = t
	})
}
