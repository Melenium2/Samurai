package retry

type RetryFunc func() error

func Retry(retryFunc RetryFunc, options ...Options) {

}