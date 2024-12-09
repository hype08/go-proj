package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LogConfig struct {
	level      zerolog.Level
	useConsole bool // flag for console vs JSON output
}

func NewLogConfig() (*LogConfig, error) {
	rawLogLevel := viper.GetString("log_level")
	if rawLogLevel == "" {
		rawLogLevel = "info"
	}

	logLevel, err := zerolog.ParseLevel(rawLogLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL: %w", err)
	}

	useConsole := viper.GetBool("log_console")
	if !viper.IsSet("log_console") {
		useConsole = true
	}

	return &LogConfig{
		level:      logLevel,
		useConsole: useConsole,
	}, nil
}

func (c *LogConfig) Apply() {
	zerolog.SetGlobalLevel(c.level)
	zerolog.DurationFieldInteger = true

	if c.useConsole {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "[15:04:05]",
			NoColor:    false,
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.MessageFieldName,
			},
		}
		log.Logger = log.Output(output)
	} else {
		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
