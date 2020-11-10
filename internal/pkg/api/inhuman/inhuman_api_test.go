package inhuman_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/api/inhuman"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type inhuman_api_mock struct {
	ExpectedCode int
	ExpectedBody interface{}
}

func (m *inhuman_api_mock) Flow(key string) ([]inhuman.App, error) {
	apps := make([]inhuman.App, 0)
	err := m.Request(m.Endpoint("mainPage"), "post", map[string]interface{} {
		"query": key,
		"hl": "13",
		"gl": "123",
		"count": 50,
	}, &apps)

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (m *inhuman_api_mock) App(bundle string) (*inhuman.App, error) {
	var app *inhuman.App
	err := m.Request(m.Endpoint("bundle"), "post", map[string]string{
		"query": bundle,
		"hl":    "en",
		"gl":    "us",
	}, &app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (m *inhuman_api_mock) Endpoint(endpoint string) string {
	return fmt.Sprintf("/%s", endpoint)
}

func (m *inhuman_api_mock) Request(endpoint, method string, data interface{}, response interface{}) error {
	var err error
	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		app := m.ExpectedBody
		b, _ := json.Marshal(app)
		w.WriteHeader(m.ExpectedCode)
		io.WriteString(w, string(b))
	}

	r := httptest.NewRequest(method, endpoint, bytes.NewReader(b))
	w := httptest.NewRecorder()
	handler(w, r)

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

/*
Successful test
============================================================================================
*/

func TestApp_ShouldMakeRequestToExternalApiAndWriteToAppResult_NoErrors(t *testing.T) {
	api := &inhuman_api_mock{
		ExpectedCode: 200,
		ExpectedBody: &inhuman.App{
			Bundle:     "123",
			Categories: "GAME",
		},
	}
	res, err := api.App("exmaple")
	assert.NoError(t, err)
	assert.Equal(t, "123", res.Bundle)
	assert.Equal(t, "GAME", res.Categories)
}

func TestEndpoint_ShouldReturnCorrectEndpointString_NoErrors(t *testing.T) {
	api := &inhuman_api_mock{}
	res := api.Endpoint("bundle")
	assert.Equal(t, "/bundle", res)
}

func TestRequest_ShouldMakeRequestToExternalApi_NoErrors(t *testing.T) {
	api := &inhuman_api_mock{}
	err := api.Request("/exmaple", "get", struct{}{}, struct{}{})
	assert.NoError(t, err)
}

func TestFlow_ShouldReturnListWithApplications_NoError(t *testing.T) {
	api := &inhuman_api_mock{
		ExpectedCode: 200,
		ExpectedBody: []inhuman.App {
			inhuman.App{ Bundle: "123"}, { Bundle: "222"},
		},
	}
	apps, err := api.Flow("car")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(apps))
}

/*
===================================================================================================
*/

type inhuman_api_mock_fail struct {
	ExpectedCode         int
	ExpectedResponseBody interface{}
}

func (m *inhuman_api_mock_fail) Flow(key string) ([]inhuman.App, error) {
	apps := make([]inhuman.App, 0)
	err := m.Request(m.Endpoint("mainPage"), "post", map[string]interface{} {
		"query": key,
		"hl": "13",
		"gl": "123",
		"count": 50,
	}, &apps)

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (m *inhuman_api_mock_fail) App(bundle string) (*inhuman.App, error) {
	return nil, m.Request(m.Endpoint("bundle"), "post", struct{}{}, struct{}{})
}

func (m *inhuman_api_mock_fail) Endpoint(endpoint string) string {
	return fmt.Sprintf("/%s", endpoint)
}

func (m *inhuman_api_mock_fail) Request(endpoint, method string, data interface{}, response interface{}) error {
	var err error
	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(m.ExpectedResponseBody)
		w.WriteHeader(m.ExpectedCode)
		io.WriteString(w, string(b))
	}

	r := httptest.NewRequest(method, endpoint, bytes.NewReader(b))
	w := httptest.NewRecorder()
	handler(w, r)

	if w.Code >= 300 {
		return errors.New("response with fail status")
	}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

/*
Unsuccessful test
============================================================================================
*/

func TestApp_ShouldReturn500ErrorOrJsonMarshalError_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode: 500,
	}
	res, err := api.App("bundle")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestEndpoint_ShouldReturn500Error_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode: 500,
		ExpectedResponseBody: map[string]string{
			"bundle": "app",
		},
	}
	err := api.Request("/bundle", "get", struct{}{}, struct{}{})
	assert.Error(t, err)
	assert.Equal(t, "response with fail status", err.Error())
}

func TestEndpoint_ShouldReturnUnpredictableBody_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode:         200,
		ExpectedResponseBody: `{ "bundle": "bundle" }]`,
	}
	err := api.Request("/bundle", "get", struct{}{}, &inhuman.App{})
	assert.Error(t, err)
	assert.Equal(t, "json: cannot unmarshal string into Go value of type inhuman.App", err.Error())
}

func TestEndpoint_ShouldReturnErrorCozIncorrectData_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode:         200,
		ExpectedResponseBody: `{ "bundle": "bundle" }`,
	}
	err := api.Request("/bundle", "get", make(chan int), &inhuman.App{})
	assert.Error(t, err)
}

func TestFlow_ShouldReturnWrongResultFromRequestIncorrectDataStruct_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode:         200,
		ExpectedResponseBody: `[{ "1": "1" }]`,
	}
	apps, err := api.Flow("car")
	assert.Error(t, err)
	assert.Nil(t, apps)
}

func TestFlow_ShouldReturnWrongResultFromRequestIncorrectDataArray_Error(t *testing.T) {
	api := &inhuman_api_mock_fail{
		ExpectedCode:         200,
		ExpectedResponseBody: `{ "1": "1" }`,
	}
	apps, err := api.Flow("car")
	assert.Error(t, err)
	assert.Nil(t, apps)
}

/*
=============================================================================
*/

func Config() inhuman.Config {
	c := config.New("../../../../config/dev.yml")

	return inhuman.Config{
		Url:       c.Api.Url,
		Key:       c.Api.Key,
		Hl:        "ru",
		Gl:        "ru",
		AppsCount: 250,
	}
}

var bundle = "com.and.wareternal"

func TestEndpoint_ShouldConcatDefaultUrlAndEndpoint_NoError(t *testing.T) {
	c := Config()
	api := inhuman.New(c)
	res := api.Endpoint("bundle")
	assert.Equal(t, c.Url + "/bundle", res)
}

func TestApp_ShouldReturnAppInformationFromApi_NoError(t *testing.T) {
	c := Config()
	api := inhuman.New(c)
	res, err := api.App(bundle)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, bundle, res.Bundle)
}

func TestFlow_ShouldReturnMainPageApps_NoError(t *testing.T) {
	c := Config()
	api := inhuman.New(c)
	res, err := api.Flow("car")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Greater(t, len(res), 0)
	assert.Greater(t, len(res),  50)
}

func TestFlow_ShouldReturnAppsFor10Keys_NoError(t *testing.T) {
	ti := time.Now()
	c := Config()
	api := inhuman.New(c)
	keys := []string {"car", "cart", "car games", "game for kids", "russian mobiles", "anime", "anime games", "wallpapers", "key", "door"}
	for _, k := range keys {
		res, err := api.Flow(k)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, len(res), 0)
		t.Log(k)
	}
	t.Log(time.Now().Sub(ti).Seconds() * 6000)
}

func TestApp_ShouldReturnErrorIfBundleIsWrong_Error(t *testing.T) {
	api := inhuman.New(Config())
	res, err := api.App("")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestApp_ShouldReturnErrorCozKeyIsIncorrect_Error(t *testing.T) {
	api := inhuman.New(Config())
	res, err := api.App("dfghadsvadkasdasdskjdsnkjdna123ad;lmsakda")
	assert.Error(t, err)
	assert.Nil(t, res)
}