package main

import (
	"encoding/json"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

type MigrationsConfig struct {
	StoragePath    string `json:"storage_path"`
	MigrationsPath string `json:"migrations_path"`
}

func main() {
	logFile, err := os.OpenFile("migrations_logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logFile.WriteString("Запуск мигратора\n")

	configFile, err := os.Open("cmd/migrator/config/migrations_config.json")
	if err != nil {
		logFile.WriteString(err.Error() + "\n")
		panic(err.Error())
	}

	var cfg MigrationsConfig
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&cfg); err != nil {
		logFile.WriteString(err.Error() + "\n")
		panic(err.Error())
	}

	m, err := migrate.New(
		"file://"+cfg.MigrationsPath,
		cfg.StoragePath,
	)
	defer m.Close()
	if err != nil {
		logFile.WriteString(err.Error() + "\n")
		panic(err.Error())
	}
	logFile.WriteString("Начало миграции\n")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logFile.WriteString(err.Error() + "\n")
		panic(err.Error())
	}
}
