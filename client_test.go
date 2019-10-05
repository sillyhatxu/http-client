package client

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const host = ""

func TestDoGet(t *testing.T) {
	httpClient := NewHttpClient(host, ShowResponseLog(true))
	response := httpClient.DoGet("")
	assert.Nil(t, response.Error)
	type Obj struct {
	}
	var obj Obj
	err := response.AnalysisBody(&obj)
	assert.Nil(t, err)
}

func TestDoPost(t *testing.T) {
	httpClient := NewHttpClient(host, ShowResponseLog(true))
	var input interface{}
	inputJSON, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	response := httpClient.DoPost("", bytes.NewBuffer(inputJSON))
	assert.Nil(t, response.Error)
	type Obj struct {
	}
	var obj Obj
	err = response.AnalysisBody(&obj)
	assert.Nil(t, err)
}
