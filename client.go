package client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

const (
	POST   = "POST"
	PUT    = "PUT"
	GET    = "GET"
	DELETE = "DELETE"
)

type Client struct {
	host   string
	config *Config
}

func NewHttpClient(host string, opts ...Option) *Client {
	//default
	config := &Config{
		header: map[string]string{
			"Content-Type": "application/json",
		},
		showResponseLog: false,
		timeout:         10 * time.Second,
		attempts:        3,
		delay:           500 * time.Millisecond,
		errorCallback: func(n uint, err error) {
			logrus.Errorf("retry [%d] request %s error. %v", n, host, err)
		},
	}
	//apply opts
	for _, opt := range opts {
		opt(config)
	}
	return &Client{
		host:   host,
		config: config,
	}
}

func (c Client) getHttpClient() *http.Client {
	return &http.Client{
		Timeout: c.config.timeout,
		//Transport: &http.Transport{
		//	Dial: (&net.Dialer{
		//		Timeout: 5 * time.Second,
		//	}).Dial,
		//	TLSHandshakeTimeout: 5 * time.Second,
		//},
	}
}

func (c Client) DoGet(url string) *Response {
	return c.do(GET, url, nil)
}

func (c Client) DoPut(url string, body io.Reader) *Response {
	return c.do(PUT, url, body)
}

func (c Client) DoPost(url string, body io.Reader) *Response {
	return c.do(POST, url, body)
}

func (c Client) DoDelete(url string) *Response {
	return c.do(DELETE, url, nil)
}

func (c Client) do(method, url string, body io.Reader) *Response {
	responseConfig := &ResponseConfig{host: c.host, url: url, method: method, showResponseLog: c.config.showResponseLog}
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.host, url), body)
	if err != nil {
		return NewResponse(nil, err, responseConfig)
	}
	for key, value := range c.config.header {
		request.Header.Set(key, value)
	}
	httpResponse, err := c.getHttpClient().Do(request)
	return NewResponse(httpResponse, err, responseConfig)
}
