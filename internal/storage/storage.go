package storage

import (
	"contest/internal/domain"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Not Found")
)

type TestRepository interface {
	AddTest(taskID int, input string, expectedResult string, points int) error
	DeleteTest(id int) error
	UpdateTest(id int, newItem domain.Test) error
	GetTests() ([]domain.Test, error)
	FindTestByID(id int) (domain.Test, error)
	FindTestsByTaskID(taskID int) ([]domain.Test, error)
	FindTestByCondition(condition func(item domain.Test) bool) (domain.Test, error)
	FindTestsByCondition(condition func(item domain.Test) bool) ([]domain.Test, error)
}

type Storage struct {
	DB             *sql.DB
	TestRepository TestRepository
}
