package config

import (
	"errors"
	"os"
)

type JWTConfig struct {
	Secret string
}

func loadJWTConfig() (*JWTConfig, error) {
	config := &JWTConfig{}
	config.Secret = os.Getenv("JWT_SECRET")
	if config.Secret == "" {
		return config, errors.New("failed to read JWT_SECRET value from the environment")
	}

	return config, nil
}
