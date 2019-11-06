package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
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
	id     string
	host   string
	config *Config
}

func NewHttpClient(host string, opts ...Option) *Client {
	//default
	config := &Config{
		header: map[string]string{
			"Content-Type": "application/json",
		},
		showResponseLog: true,
		timeout:         10 * time.Second,
	}
	for _, opt := range opts {
		opt(config)
	}
	id := GetId()
	logrus.Infof("create http client {host : %s, id : %s}", host, id)
	return &Client{
		id:     id,
		host:   host,
		config: config,
	}
}

func (c *Client) getHttpClient() *http.Client {
	return &http.Client{
		//Timeout: c.config.timeout,
	}
}

func (c *Client) DoGet(url string) (*http.Response, error) {
	return c.do(GET, url, nil)
}

func (c *Client) DoPut(url string, body io.Reader) (*http.Response, error) {
	return c.do(PUT, url, body)
}

func (c *Client) DoPost(url string, body io.Reader) (*http.Response, error) {
	return c.do(POST, url, body)
}

func (c *Client) DoDelete(url string) (*http.Response, error) {
	return c.do(DELETE, url, nil)
}

type ChannelResponse struct {
	Response *http.Response
	Err      error
}

func (c *Client) do(method, url string, body io.Reader) (*http.Response, error) {
	logrus.Infof("http[%s] request %s %s%s. body : %v", c.id, method, c.host, url, body)
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.host, url), body)
	if err != nil {
		return nil, err
	}
	for key, value := range c.config.header {
		request.Header.Set(key, value)
	}

	customChannel := make(chan ChannelResponse)
	d := time.Now().Add(c.config.timeout)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	go func() {
		res, err := c.getHttpClient().Do(request)
		customChannel <- ChannelResponse{Response: res, Err: err}
	}()
	select {
	case channelResponse := <-customChannel:
		return channelResponse.Response, channelResponse.Err
	case <-ctx.Done():
		return nil, TimeOut
	}
}

func (c *Client) AnalysisBody(res *http.Response, input interface{}) error {
	if res == nil {
		logrus.Errorf("get http code error. http response is nil")
		return fmt.Errorf("get http code error. httpResponse is nil")
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("analysis response body error. %v", err)
		return err
	}
	if c.config.showResponseLog {
		logrus.Infof("http[%s] response : %s", c.id, string(bodyBytes))
	}
	if err := json.Unmarshal(bodyBytes, input); err != nil {
		return err
	}
	return nil
}
