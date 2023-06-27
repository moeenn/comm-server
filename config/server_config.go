package config

type ServerConfig struct {
	Host string
	Port string
}

func (conf *ServerConfig) HostPort() string {
	return conf.Host + conf.Port
}

func loadServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: "0.0.0.0",
		Port: ":5000",
	}
}
