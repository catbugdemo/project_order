package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpGet(t *testing.T) {
	type Te struct {
		Test string `json:"test"`
		Name string `json:"name"`
	}
	var te Te
	err := HttpGet("http://localhost:9098/test", &te)
	assert.Nil(t, err)
}

func TestInitSms(t *testing.T) {
	InitSms()
}
