package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var c = Config{
	taskURL:   "url",
	taskToken: "token",
}

func TestConfig_GetTaskToken(t *testing.T) {
	t.Run("Should return the token", func(t *testing.T) {
		assert.Equal(t, "token", c.GetTaskToken())
	})
}

func TestConfig_GetTaskURL(t *testing.T) {
	t.Run("Should return the url", func(t *testing.T) {
		assert.Equal(t, "url", c.GetTaskURL())
	})
}

func TestGetConfig(t *testing.T) {
	t.Run("Should return an error when TASK_URL is not set", func(t *testing.T) {
		_, err := GetConfig()
		assert.NotNil(t, err)
	})
	assert.Nil(t, os.Setenv(envTaskURL, "url"))

	t.Run("Should return an error when TASK_TOKEN is not set", func(t *testing.T) {
		_, err := GetConfig()
		assert.NotNil(t, err)
	})

	assert.Nil(t, os.Setenv(envTaskToken, "token"))

	t.Run("Should return an config with no error", func(t *testing.T) {
		conf, err := GetConfig()
		assert.Nil(t, err)
		assert.Equal(t, "url", conf.GetTaskURL())
		assert.Equal(t, "token", conf.GetTaskToken())
	})
}
