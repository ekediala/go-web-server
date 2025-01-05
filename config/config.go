package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port  string `validate:"required,min=2"`
	Admin bool   `validate:"required"`
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

	err := v.Struct(v)
	if err != nil {
		return Config{}, fmt.Errorf("validating config:%w", err)
	}

	return config, nil
}
