package config

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

func setupViper() {
	viper.SetDefault("log_level", "info")
	viper.AutomaticEnv()
}

func getDatabaseUrl() (string, error) {
	databaseUrl := viper.GetString("database_url")

	if databaseUrl == "" {
		return "", errors.New("DATABASE_URL is required.")
	}

	parsedUrl, err := url.Parse(databaseUrl)
	if err != nil {
		return "", fmt.Errorf("Invalid Database URL: %w", err)
	}

	query := parsedUrl.Query()          /* parse: params into key-value map */
	query.Set("sslmode", "disable")     /* modify: unencrypted plain text till security stuff */
	parsedUrl.RawQuery = query.Encode() /* store: URL encode all parameters */

	return parsedUrl.String(), nil
}
