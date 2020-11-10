package inhuman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Interface who provide api to external api
type ExternalApi interface {
	App(bundle string) (*App, error)
	Flow(key string) ([]App, error)
}

// Struct of external api instance
type inhumanApi struct {
	config Config
}

// Get application info by bundle from external api
func (api *inhumanApi) App(bundle string) (*App, error) {
	var app *App
	err := api.Request(api.Endpoint("bundle"), "post", map[string]string{
		"query": bundle,
		"hl": api.config.Hl,
		"gl": api.config.Gl,
	}, &app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

// Method calls api and return top N application from main page
func (api *inhumanApi) Flow(key string) ([]App, error) {
	apps := make([]App, 0)
	err := api.Request(api.Endpoint("mainPage"), "post", map[string]interface{} {
		"query": key,
		"hl": api.config.Hl,
		"gl": api.config.Gl,
		"count": api.config.AppsCount,
	}, &apps)

	if err != nil {
		return nil, err
	}

	return apps, nil
}

// Make request to API to the given endpoint and return data to the pointer
// or error
func (api *inhumanApi) Request(endpoint, method string, data interface{}, response interface{}) error {
	var err error
	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(strings.ToUpper(method), endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", api.config.Key)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 200 {
		return fmt.Errorf("external api response with status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

// Generate full url to api from string
func (api *inhumanApi) Endpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", api.config.Url, endpoint)
}

// Create new instance of inhuman api
func New(config Config) *inhumanApi {
	return &inhumanApi{
		config: config,
	}
}

