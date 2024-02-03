package postgres

import (
	. "contest/internal/domain"
	"contest/internal/storage"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

type TestRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewTestRepository(db *sql.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (r *TestRepository) AddTest(taskID int, input string, expectedResult string, points int) error {
	_, err := r.db.Exec("INSERT INTO tests(task_id, input, expected_result, points) VALUES ($1, $2, $3, $4)",
		taskID, input, expectedResult, points)
	if err != nil {
		err = fmt.Errorf("In TestRepository(AddItem): %w", err)
	}
	return err
}

func (r *TestRepository) DeleteTest(id int) error {
	rows, err := r.db.Exec("DELETE from tests where id=$1", id)
	if err != nil {
		err = fmt.Errorf("In TestRepository(DeleteItem): %w", err)
		return err
	}
	if count, err := rows.RowsAffected(); count == 0 {
		if err != nil {
			err = fmt.Errorf("In TestRepository(DeleteItem): %w", err)
			return err
		}
		return storage.ErrNotFound
	}
	return nil
}

func (r *TestRepository) UpdateTest(id int, newItem Test) error {
	_, err := r.db.Exec("UPDATE tests SET id=$1,task_id=$2, input=$3, expected_result=$4, points=$5 WHERE id=$6",
		newItem.ID, newItem.TaskID, newItem.Input, newItem.ExpectedResult, newItem.Points, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrNotFound
		}
		err = fmt.Errorf("In TestRepository(UpdateItem): %w", err)
		return err
	}
	return nil
}

func (r *TestRepository) GetTests() ([]Test, error) {
	rows, err := r.db.Query("SELECT id, task_id, input, expected_result, points FROM tests")
	if err != nil {
		return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var test Test
		err = rows.Scan(&test.ID, &test.TaskID, &test.Input, &test.ExpectedResult, &test.Points)
		if err != nil {
			return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
		}
		tests = append(tests, test)
	}
	return tests, nil
}

func (r *TestRepository) FindTestByID(id int) (Test, error) {
	row := r.db.QueryRow("SELECT id, task_id, input, expected_result, points FROM tests WHERE id = $1", id)
	var test Test
	err := row.Scan(&test.ID, &test.TaskID, &test.Input, &test.ExpectedResult, &test.Points)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}
		err = fmt.Errorf("In TestRepository(FindItemByID): %w", err)
	}
	return test, err
}

func (r *TestRepository) FindTestsByTaskID(taskID int) ([]Test, error) {
	rows, err := r.db.Query("SELECT id, task_id, input, expected_result, points FROM tests WHERE task_id = $1", taskID)
	if err != nil {
		return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var test Test
		err = rows.Scan(&test.ID, &test.TaskID, &test.Input, &test.ExpectedResult, &test.Points)
		if err != nil {
			return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
		}
		tests = append(tests, test)
	}
	return tests, nil
}

func (r *TestRepository) FindTestByCondition(condition func(item Test) bool) (Test, error) {
	items, err := r.FindTestsByCondition(condition)
	if err != nil {
		return Test{}, fmt.Errorf("In TestRepository(FindItemByCondition): %w", err)
	}
	if len(items) == 0 {
		return Test{}, storage.ErrNotFound
	}
	return items[0], nil
}

func (r *TestRepository) FindTestsByCondition(condition func(item Test) bool) ([]Test, error) {
	table, err := r.GetTests()
	if err != nil {
		return nil, fmt.Errorf("In TestRepository(FindItemsByCondition): %w", err)
	}
	var res []Test
	for _, test := range table {
		if condition(test) {
			res = append(res, test)
		}
	}
	return res, nil
}
