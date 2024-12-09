package config

import (
	"errors"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	port        int
	databaseUrl string
	log         *LogConfig
}

func NewServerConfig() (*ServerConfig, error) {
	setupViper()
	port := viper.GetInt("port")

	if port == 0 {
		return nil, errors.New("PORT is required.")
	}

	databaseUrl, err := getDatabaseUrl()
	if err != nil {
		return nil, err
	}

	log, err := NewLogConfig()
	if err != nil {
		return nil, err
	}
	return &ServerConfig{
		port,
		databaseUrl,
		log,
	}, nil
}

func (c *ServerConfig) DatabaseUrl() string {
	return c.databaseUrl
}

func (c *ServerConfig) Port() int {
	return c.port
}

func (c *ServerConfig) Log() *LogConfig {
	return c.log
}
