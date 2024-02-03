package main

import (
	"contest/internal/config"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	log.Println("Начало миграции")
	var configPath string
	flag.StringVar(&configPath, "path", "", "config path")
	flag.Parse()
	if configPath == "" {
		log.Fatal("Необходимо указать путь до файлов конфигурации")
	}
	cfg := config.MustLoadMigrationConfig(configPath)

	m, err := migrate.New(
		"file://"+cfg.MigrationsPath,
		cfg.StoragePath,
	)
	defer m.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Начало миграции")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("Завершение миграции")
}
