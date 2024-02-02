package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

var (
	localEnv = "local"
	devEnv   = "dev"
)

type Config struct {
	Port             int    `env:"PORT"`
	ApiKey           string `env:"API_KEY"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	Env              string `env:"ENV"`
}

func MustLoadAppConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	cfg := &Config{}

	err := envconfig.Process(context.Background(), cfg)
	if err != nil {
		panic("Load Config error: " + err.Error())
	}

	return cfg
}
