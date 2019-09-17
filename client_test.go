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
	response.AnalysisBody()
}
