package config

import (
	"errors"
	"os"
)

type DatabaseConfig struct {
	URI string
}

func loadDatabaseConfig() (*DatabaseConfig, error) {
	config := &DatabaseConfig{}
	uri := os.Getenv("DATABASE_URI")

	if uri == "" {
		return config, errors.New("failed to read DATABASE_URI from environment")
	}

	config.URI = uri
	return config, nil
}
