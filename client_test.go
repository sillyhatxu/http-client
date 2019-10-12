package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
	"time"
)

const host = "http://localhost:8080/user-backend"

type UserResponse struct {
	Code string `json:"code"`
	Data *User  `json:"data"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TestDoGet(t *testing.T) {
	httpClient := NewHttpClient(host)
	response, err := httpClient.DoGet(fmt.Sprintf("/users/%s", "TEST_001"))
	if err != nil {
		panic(err)
	}
	var res UserResponse
	err = httpClient.AnalysisBody(response, &res)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, res.Data.Id, "TEST_001")
	assert.EqualValues(t, res.Data.Name, "silly hat")
}

func TestTimeOut(t *testing.T) {
	httpClient := NewHttpClient(host, Header(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}), Timeout(1*time.Millisecond))
	response, err := httpClient.DoGet(fmt.Sprintf("/users/%s", "TEST_001"))
	assert.Nil(t, response)
	fmt.Println(err.Error())
}

func TestDoPost(t *testing.T) {
	httpClient := NewHttpClient(host, Header(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}), Timeout(500*time.Millisecond))
	data := url.Values{}
	data.Set("params1", "value1")
	data.Add("params2", "value2-1")
	data.Add("params2", "value2-2")
	response, err := httpClient.DoPost("/users", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	var res interface{}
	err = httpClient.AnalysisBody(response, &res)
	assert.Nil(t, err)
	fmt.Println(fmt.Sprintf("%#v", res))
}
