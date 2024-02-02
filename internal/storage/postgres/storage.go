package postgres

import (
	"contest/internal/storage"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewStorage(user string, password string, dbName string) (*storage.Storage, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", user, password, dbName))
	if err != nil {
		return nil, err
	}

	testRepository := NewTestRepository(db)

	return &storage.Storage{
		DB:             db,
		TestRepository: testRepository,
	}, nil
}
