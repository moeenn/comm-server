package config

import (
	"errors"
	"os"
)

type JWTConfig struct {
	ServerSecret string
	ClientSecret string
}

func loadJWTConfig() (*JWTConfig, error) {
	config := &JWTConfig{}

	config.ServerSecret = os.Getenv("JWT_SERVER_SECRET")
	if config.ServerSecret == "" {
		return config, errors.New("failed to read JWT_SERVER_SECRET value from the environment")
	}

	config.ClientSecret = os.Getenv("JWT_CLIENT_SECRET")
	if config.ClientSecret == "" {
		return config, errors.New("failed to read JWT_CLIENT_SECRET value from the environment")
	}

	return config, nil
}
