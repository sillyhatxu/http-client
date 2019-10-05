package client

import (
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
