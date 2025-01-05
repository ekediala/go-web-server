package config

import (
	"os"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string `validate:"required,min=4"`
}

type DBConfig struct {
	URL string `validate:"required"`
}

func LoadConfig() (Config, error) {
	server := ServerConfig{
		Port: os.Getenv("PORT"),
	}

	db := DBConfig{
		URL: os.Getenv("DB_URL"),
	}

	config := Config{
		Server: server,
		DB:     db,
	}

	v := validator.New(validator.WithRequiredStructEnabled())

	err := v.Struct(config)

	return config, err
}
