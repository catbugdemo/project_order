package inits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
