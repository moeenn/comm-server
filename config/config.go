package config

type Config struct {
	JWTConfig      *JWTConfig
	DatabaseConfig *DatabaseConfig
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	jwtConfig, err := loadJWTConfig()
	if err != nil {
		return config, err
	}

	databaseConfig, err := loadDatabaseConfig()
	if err != nil {
		return config, err
	}

	config.JWTConfig = jwtConfig
	config.DatabaseConfig = databaseConfig

	return config, nil
}
