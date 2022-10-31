package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	Level      string // уровень логирования
	LoggerName string // имя используемого логера
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path ...string) (config *Config, err error) {

	if err := godotenv.Load(path...); err != nil {
		return nil, err
	}
	cfg := &Config{
		Level:      os.Getenv("LOG_LEVEL"),
		LoggerName: os.Getenv("LOGGER_NAME"),
	}

	return cfg, nil
}
