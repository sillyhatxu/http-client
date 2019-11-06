package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const host = "http://localhost:8080"

type UserResponse struct {
	Code string `json:"code"`
	Data *User  `json:"data"`
}

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobileNumber"`
}

func TestDoGet(t *testing.T) {
	httpClient := NewHttpClient(host)
	response, err := httpClient.DoGet(fmt.Sprintf("/user-info/%s", "U5B5FC45D564F127D897B9258"))
	if err != nil {
		panic(err)
	}
	var res UserResponse
	err = httpClient.AnalysisBody(response, &res)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, res.Data.Id, "U5B5FC45D564F127D897B9258")
	assert.EqualValues(t, res.Data.Name, "123")
	assert.EqualValues(t, res.Data.MobileNumber, "+86176880808888")
}

func TestTimeOut(t *testing.T) {
	httpClient := NewHttpClient(host, Header(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}), Timeout(1*time.Millisecond))
	response, err := httpClient.DoGet(fmt.Sprintf("/user-info/%s", "U5B5FC45D564F127D897B9258"))
	assert.Nil(t, response)
	assert.EqualValues(t, err, TimeOut)
}
