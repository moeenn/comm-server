package config

import (
	"errors"
	"os"
)

var (
	errDBUriRead = errors.New("failed to read DATABASE_URI from environment")
)

type DatabaseConfig struct {
	URI string
}

func loadDatabaseConfig() (*DatabaseConfig, error) {
	config := &DatabaseConfig{}
	uri := os.Getenv("DATABASE_URI")

	if uri == "" {
		return config, errDBUriRead
	}

	config.URI = uri
	return config, nil
}
