package inhuman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)


// Make request to API to the given endpoint and return data to the pointer
// or error
func Request(endpoint, method string, options ...RequestOption) error {
	opt := defaultOptions()
	for _, o := range options {
		o.apply(&opt)
	}
	var err error
	var r io.Reader
	if opt.data != nil {
		b, err := json.Marshal(opt.data)
		if err != nil {
			return err
		}
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequest(strings.ToUpper(method), Endpoint(opt.url, endpoint), r)
	if err != nil {
		return err
	}

	if opt.query != nil {
		q := req.URL.Query()
		for k, v := range opt.query {
			q.Add(k, fmt.Sprint(v))
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	req.Header.Set("Content-Type",  "application/json")
	req.Header.Set("Authorization", opt.apikey)


	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("external api response with status %d, message = %s", resp.StatusCode, string(b))
	}

	if err := json.NewDecoder(resp.Body).Decode(&opt.response); err != nil {
		return err
	}

	return nil
}

// Generate full url to api from string
func  Endpoint(url, endpoint string) string {
	return fmt.Sprintf("%s/%s", url, endpoint)
}
