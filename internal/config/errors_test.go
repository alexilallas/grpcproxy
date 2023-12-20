package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingEnv_Error(t *testing.T) {
	err := MissingEnv("ENV_MISSING")
	expectedMessage := "missing env ENV_MISSING"
	t.Run("Should return an error message", func(t *testing.T) {
		assert.Equal(t, expectedMessage, err.Error())
	})
}
