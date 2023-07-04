package config

import (
	"errors"
	"os"
)

var (
	errServerSecretRead = errors.New("failed to read JWT_SERVER_SECRET value from the environment")
	errClientSecretRead = errors.New("failed to read JWT_CLIENT_SECRET value from the environment")
)

type JWTConfig struct {
	ServerSecret string
	ClientSecret string
}

func loadJWTConfig() (*JWTConfig, error) {
	config := &JWTConfig{}

	config.ServerSecret = os.Getenv("JWT_SERVER_SECRET")
	if config.ServerSecret == "" {
		return config, errServerSecretRead
	}

	config.ClientSecret = os.Getenv("JWT_CLIENT_SECRET")
	if config.ClientSecret == "" {
		return config, errClientSecretRead
	}

	return config, nil
}
