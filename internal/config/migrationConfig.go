package config

import (
	"encoding/json"
	"os"
)

type MigrationsConfig struct {
	StoragePath    string `json:"storage_path"`
	MigrationsPath string `json:"migrations_path"`
}

func MustLoadMigrationConfig(configPath string) MigrationsConfig {
	configFile, err := os.Open(configPath)

	if err != nil {
		panic(err.Error())
	}

	var cfg MigrationsConfig
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
