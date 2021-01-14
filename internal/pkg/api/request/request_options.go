package request

// requestOptions struct contains request options
type requestOptions struct {
	data interface{}
	response interface{}
	apikey string
	url string
	query map[string]interface{}
}

// interface represent single RequestOption
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

// WithData added data to request to be passed
func WithData(data interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.data = data
	})
}

// WithResponseType contains the type of struct that will return
func WithResponseType(resp interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.response = resp
	})
}

// WithApikey contains access token to api
func WithApikey(key string) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.apikey = key
	})
}

// WithUrl contains url for request
func WithUrl(url string) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.url = url
	})
}

// WithQueryParams contains map of k:v to be passed as query params
func WithQueryParams(params map[string]interface{}) RequestOption {
	return newFuncRequestOptions(func(o *requestOptions) {
		o.query = params
	})
}


func defaultOptions() requestOptions {
	return requestOptions{}
}
