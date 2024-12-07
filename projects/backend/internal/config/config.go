package config

import "github.com/spf13/viper"

func setupViper() {
	viper.SetDefault("log_level", "info")
	viper.AutomaticEnv()
}

func getDatabaseUrl() (string, error) {
}
