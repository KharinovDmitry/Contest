package storage

import (
	"contest/internal/domain"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Not Found")
)

type Repository[T any] interface {
	AddItem(taskID int, input string, expectedResult string, points int) error
	DeleteItem(id int) error
	UpdateItem(id int, newItem T) error
	GetTable() ([]T, error)
	FindItemByID(id int) (T, error)
	FindItemByCondition(condition func(item T) bool) (T, error)
	FindItemsByCondition(condition func(item T) bool) ([]T, error)
}

type Storage struct {
	DB             *sql.DB
	TestRepository Repository[domain.Test]
}
