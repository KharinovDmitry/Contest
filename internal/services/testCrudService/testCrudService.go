package testCrudService

import (
	. "contest/internal/domain"
	"contest/internal/storage"
	"errors"
	"fmt"
	"log/slog"
)

var (
	TestsNotFoundError = errors.New("Tests not found")
	ErrNotFound        = errors.New("Not found in database")
)

type TestCrudService struct {
	logger         *slog.Logger
	testRepository storage.TestRepository
}

func NewTestCrudService(testRepository storage.TestRepository) *TestCrudService {
	return &TestCrudService{
		testRepository: testRepository,
	}
}

func (s *TestCrudService) GetTest(id int) (Test, error) {
	test, err := s.testRepository.FindTestByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return Test{}, ErrNotFound
		}
		return Test{}, fmt.Errorf("In TestService(GetTest): %w", err)
	}
	return test, nil
}

func (s *TestCrudService) AddTest(taskID int, input string, expectedResult string, points int) error {
	err := s.testRepository.AddTest(taskID, input, expectedResult, points)
	if err != nil {
		return fmt.Errorf("In TestService(AddTest): %w", err)
	}
	return nil
}

func (s *TestCrudService) DeleteTest(id int) error {
	err := s.testRepository.DeleteTest(id)
	if err != nil {
		return fmt.Errorf("In TestService(DeleteTest): %w", err)
	}
	return nil
}

func (s *TestCrudService) UpdateTest(id int, newTest Test) error {
	err := s.testRepository.UpdateTest(id, newTest)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("In TestService(UpdateTest): %w", err)
	}
	return nil
}

func (s *TestCrudService) GetTests() ([]Test, error) {
	tests, err := s.testRepository.GetTests()
	if err != nil {
		return nil, fmt.Errorf("In TestService(GetTests): %w", err)
	}
	return tests, nil
}

func (s *TestCrudService) GetTestsByTaskID(taskID int) ([]Test, error) {
	tests, err := s.testRepository.FindTestsByTaskID(taskID)
	if err != nil {
		return nil, fmt.Errorf("In TestService(GetTestsByTaskID): %w", err)
	}
	return tests, nil
}
