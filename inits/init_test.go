package inits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("GetConfig", func(t *testing.T) {
		getConfig := GetConfig()
		get := getConfig.Get("database.server")
		fmt.Println(get)
	})
}

func TestDB(t *testing.T) {
	t.Run("DB", func(t *testing.T) {
		g := DB()
		assert.NotNil(t, g)
	})
}

func TestRedis(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		get := Get()
		assert.NotNil(t, get)
	})
}
