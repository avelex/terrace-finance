package config

import (
	"fmt"
	"os"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DebugLevel = "debug"

type Config struct {
	LogLevel      string    `env:"LOG_LEVEL" env-default:"info"`
	EventscaleURL string    `env:"EVENTSCALE_URL" env-required:"true"`
	DB            string    `env:"DB_URL" env-required:"true"`
	PrivateKey    string    `env:"PRIVATE_KEY" env-required:"true"`
	Protocol      Protocol  `yaml:"protocol"`
	Networks      []Network `yaml:"networks"`
}

type Protocol struct {
	Hub        enum.CircleDomain `yaml:"hub"`
	Stablecoin common.Address    `yaml:"stablecoin"`
	Staking    common.Address    `yaml:"staking"`
}

type Network struct {
	Name   string         `yaml:"name"`
	Domain uint32         `yaml:"domain"`
	URL    string         `yaml:"rpc_url"`
	Vault  common.Address `yaml:"vault"`
}

func Load(path string) (Config, error) {
	config := Config{}
	if err := cleanenv.ReadEnv(&config); err != nil {
		return Config{}, fmt.Errorf("read env: %w", err)
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel := zerolog.InfoLevel
	if config.LogLevel == DebugLevel {
		logLevel = zerolog.DebugLevel
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Level(logLevel)

	return config, nil
}
