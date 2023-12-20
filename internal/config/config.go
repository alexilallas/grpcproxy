package config

import "os"

const (
	envTaskURL   = "TASK_URL"
	envTaskToken = "TASK_TOKEN"
)

type Config struct {
	taskURL   string
	taskToken string
}

func GetConfig() (Config, error) {
	url, token := os.Getenv(envTaskURL), os.Getenv(envTaskToken)
	if url == "" {
		return Config{}, MissingEnv(envTaskURL)
	}
	if token == "" {
		return Config{}, MissingEnv(envTaskToken)
	}
	return Config{taskURL: url, taskToken: token}, nil
}

func (c Config) GetTaskURL() string {
	return c.taskURL
}

func (c Config) GetTaskToken() string {
	return c.taskToken
}
