package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DebugLevel = "debug"

type Config struct {
	LogLevel      string `env:"LOG_LEVEL" env-default:"info"`
	EventscaleURL string `env:"EVENTSCALE_URL" env-required:"true"`
	DB            string `env:"DB_URL" env-required:"true"`
}

func Load() (Config, error) {
	config := Config{}
	if err := cleanenv.ReadEnv(&config); err != nil {
		return Config{}, fmt.Errorf("read env: %w", err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel := zerolog.InfoLevel
	if config.LogLevel == DebugLevel {
		logLevel = zerolog.DebugLevel
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Level(logLevel)

	return config, nil
}
