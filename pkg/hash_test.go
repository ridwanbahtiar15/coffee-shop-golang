package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pwd = "123456"
var hashed string
var hc = InitHashConfig().UseDefaultConfig()

func TestGenHashedPassword(t *testing.T) {
	hp, err := hc.GenHashedPassword(pwd)
	assert.NoError(t, err, "Error During Hashing")
	assert.NotEqual(t, pwd, hp, "Failed hash password")
	hashed = hp
}


func TestComparePasswordAndHash(t *testing.T) {
	t.Run("Case password valid", func(t *testing.T) {
		valid, err := hc.ComparePasswordAndHash(pwd, hashed)
		assert.NoError(t, err, "Error during comparation")
		assert.True(t, valid, "Failed comparation")
	})

	t.Run("Case password invalid", func (t *testing.T) {
		anotherPwd := "54321"
		valid, err := hc.ComparePasswordAndHash(anotherPwd, hashed)
		assert.NoError(t, err, "Error during comparation")
		assert.False(t, valid, "Failed comparation")
	})
}