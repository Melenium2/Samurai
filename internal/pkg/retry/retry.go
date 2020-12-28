package retry

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type RetryFunc func() error

func Go(retryFunc RetryFunc, options ...Option) error {
	opt := defaultOptions()

	for _, o := range options {
		o.apply(&opt)
	}

	var errorsLog Error
	for !opt.backoff.IsEnd()  {
		err := retryFunc()

		if err != nil {
			errorsLog = append(errorsLog, err)

			delay := opt.backoff.NextDelay()
			log.Print(fmt.Sprintf("wait a %.1f secs for each iteration", delay.Seconds()))

			select {
			case <-time.After(delay):
			case <-opt.ctx.Done():
				return opt.ctx.Err()
			}
		} else {
			return nil
		}

		opt.backoff.Next()
	}

	return errorsLog
}

type Error []error

func (e Error) Error() string {
	ers := make([]string, len(e))
	for i, er := range e {
		ers[i] = er.Error()
	}
	return strings.Join(ers, "\n")
}



