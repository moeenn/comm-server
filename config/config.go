package config

type Config struct {
	ServerConfig   *ServerConfig
	DatabaseConfig *DatabaseConfig
	JWTConfig      *JWTConfig
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	config.ServerConfig = loadServerConfig()
	databaseConfig, err := loadDatabaseConfig()
	if err != nil {
		return config, err
	}

	jwtConfig, err := loadJWTConfig()
	if err != nil {
		return config, err
	}

	config.DatabaseConfig = databaseConfig
	config.JWTConfig = jwtConfig

	return config, nil
}
