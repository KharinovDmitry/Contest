package runTestService

import (
	"contest/internal/domain"
	"contest/internal/executors"
	"contest/internal/services/testCrudService"
	"contest/internal/storage"
	"errors"
	"fmt"
	"strings"
)

type IExecutor interface {
	Execute(input string, memoryLimitInKb int, timeLimitInMs int) (output string, err error)
	Close() error
}

type IExecutorFactory interface {
	NewExecutor(code string, language domain.Language) (IExecutor, error)
}

type RunTestService struct {
	codeRunnerFactory IExecutorFactory
	testRepository    storage.TestRepository
}

func NewRunTestService(codeRunnerFactory IExecutorFactory, testRepository storage.TestRepository) *RunTestService {
	return &RunTestService{
		codeRunnerFactory: codeRunnerFactory,
		testRepository:    testRepository,
	}
}

func (s *RunTestService) RunTest(taskID int, language domain.Language, code string, memoryLimitInKb int, timeLimitInMs int) (domain.TestsResult, error) {
	program, err := s.codeRunnerFactory.NewExecutor(code, language)
	if err != nil {
		return domain.TestsResult{}, fmt.Errorf("In RunTestService(RunTest): %w", err)
	}
	defer program.Close()

	tests, err := s.testRepository.FindTestsByTaskID(taskID)
	if err != nil {
		return domain.TestsResult{}, fmt.Errorf("In RunTestService(RunTest): %w", err)
	}

	if len(tests) == 0 {
		return domain.TestsResult{}, testCrudService.TestsNotFoundError
	}

	var points int
	for _, test := range tests {
		output, err := program.Execute(test.Input, memoryLimitInKb, timeLimitInMs)
		if err != nil {
			if errors.Is(err, executors.TimeLimitError) {
				return domain.TestsResult{
					ResultCode: domain.TimeLimitCode,
					Points:     points,
				}, nil
			}
			if errors.Is(err, executors.RuntimeError) {
				return domain.TestsResult{
					ResultCode:  domain.RuntimeErrorCode,
					Description: fmt.Sprintf("Error Info: %s Output: %s", err.Error(), output),
					Points:      points,
				}, nil
			}
			return domain.TestsResult{}, fmt.Errorf("In TestService(RunTests): %w", err)
		}
		if strings.Replace(output, "\n", "", 1) != test.ExpectedResult {
			return domain.TestsResult{
				ResultCode:  domain.IncorrectAnswerCode,
				Description: fmt.Sprintf("Test Failed: %d Expected: %s Actual: %s", test.ID, test.ExpectedResult, output),
				Points:      points,
			}, nil
		}
		points += test.Points
	}

	return domain.TestsResult{
		ResultCode: domain.SuccesCode,
		Points:     points,
	}, err
}
