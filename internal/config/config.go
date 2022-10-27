package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	// уровень логирования
	Level      string
	// имя используемого логера
	LoggerName string
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
