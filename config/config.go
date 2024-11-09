package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT" env-default:"7777" env-description:"Порт сервера"`
		Host string `env:"SERVER_HOST" env-default:"localhost" env-description:"Хост сервера"`
	}
	Database struct {
		Port string `env:"DB_PORT" env-default:"7777" env-description:"Порт базы данных"`
		Host string `env:"DB_HOST" env-default:"localhost" env-description:"Хост базы данных"`
	}
}

func MustNew() *Config {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFilePath := filepath.Join(workingDir, ".env")

	var cfg Config
	err = cleanenv.ReadConfig(envFilePath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
