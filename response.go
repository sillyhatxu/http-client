package client

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Response struct {
	HttpResponse   *http.Response
	Error          error
	ResponseConfig *ResponseConfig
}

func NewResponse(HttpResponse *http.Response, Error error, ResponseConfig *ResponseConfig) *Response {
	return &Response{
		HttpResponse:   HttpResponse,
		Error:          Error,
		ResponseConfig: ResponseConfig,
	}
}

type ResponseConfig struct {
	host            string
	url             string
	method          string
	showResponseLog bool
}

func (r Response) GetHttpCode() int {
	if r.HttpResponse == nil {
		return http.StatusInternalServerError
	}
	return r.HttpResponse.StatusCode
}

func (r Response) AnalysisBody(input interface{}) error {
	if r.HttpResponse == nil {
		logrus.Errorf("get http code error. httpResponse is nil. %v", r.Error)
		return fmt.Errorf("get http code error. httpResponse is nil. %v", r.Error)
	}
	bodyBytes, err := ioutil.ReadAll(r.HttpResponse.Body)
	if err != nil {
		logrus.Errorf("analysis response body error. %v", err)
		return err
	}
	if r.ResponseConfig.showResponseLog {
		logrus.Infof("Response [%s]%s%s : %s", r.ResponseConfig.method, r.ResponseConfig.host, r.ResponseConfig.url, string(bodyBytes))
	}
	if err := json.Unmarshal(bodyBytes, input); err != nil {
		return err
	}
	return nil
}
