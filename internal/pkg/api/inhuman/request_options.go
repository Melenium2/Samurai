package inhuman

type requestOptions struct {
	data interface{}
	response interface{}
	apikey string
	url string
	query map[string]interface{}
}

type RequestOption interface {
	apply(option *requestOptions)
}

type funcRequestOptions struct {
	f func(option *requestOptions)
}

func (fro *funcRequestOptions) apply(option *requestOptions) {
	fro.f(option)
}

func newFuncRequestOptions(f func(o *requestOptions)) *funcRequestOptions {
	return &funcRequestOptions{
		f: f,
	}
}

func WithData(data interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.data = data
	})
}

func WithResponseType(resp interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.response = resp
	})
}

func WithApikey(key string) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.apikey = key
	})
}

func WithUrl(url string) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.url = url
	})
}

func WithQueryParams(params map[string]interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.query = params
	})
}

func defaultOptions() requestOptions {
	return requestOptions{}
}
