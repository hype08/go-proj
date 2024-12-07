package config

import (
	"errors"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	port        int
	databaseUrl string
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
	return &ServerConfig{
		port,
		databaseUrl,
	}, nil
}
