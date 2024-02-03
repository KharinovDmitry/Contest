package services

import (
	"contest/internal/compiler"
	. "contest/internal/domain"
	"contest/internal/storage"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const (
	memoryLimitInKB = 1024
	timeLimit       = 1 * time.Second
)

var (
	ProgramError         = errors.New("Program error")
	TimeLimitError       = errors.New("Time limit error")
	UnknownLanguageError = errors.New("Unknown language")
	TestsNotFoundError   = errors.New("Tests not found")
	ErrNotFound          = errors.New("Not found in database")
)

type ITestService interface {
	RunTest(taskID int, language Language, code string) (TestsResult, error)
	GetTest(id int) (Test, error)
	AddTest(taskID int, input string, expectedResult string, points int) error
	DeleteTest(id int) error
	UpdateTest(id int, newTest Test) error
	GetTests() ([]Test, error)
	GetTestsByTaskID(taskID int) ([]Test, error)
}

type TestService struct {
	logger         *slog.Logger
	compileService compiler.Compiler
	testRepository storage.TestRepository
}

func NewTestService(compileService compiler.Compiler, testRepository storage.TestRepository) *TestService {
	return &TestService{
		compileService: compileService,
		testRepository: testRepository,
	}
}

func (s *TestService) RunTest(taskID int, language Language, code string) (TestsResult, error) {
	var fileName string
	var err error

	switch language {
	case CPP:
		fileName, err = s.compileService.CompileCPP(code)
	case Python:
		fileName, err = s.compileService.CompilePython(code)
	default:
		return TestsResult{}, fmt.Errorf("%w: %s", UnknownLanguageError, language)
	}
	if err != nil {
		return TestsResult{}, fmt.Errorf("In TestService(RunTest): %w", err)
	}

	return s.RunTestOnFile(fileName, taskID, true)
}

func (s *TestService) RunTestOnFile(fileName string, taskID int, deleteAfter bool) (TestsResult, error) {
	file, err := os.Open(fileName)
	if deleteAfter {
		defer func() {
			err := os.Remove(fileName)
			if err != nil {
				s.logger.Error(fmt.Sprintf("File not deleted: " + err.Error()))
			}
		}()
	}
	defer file.Close()
	if err != nil {
		return TestsResult{}, fmt.Errorf("In TestService(RunTestOnFile): %w", err)
	}

	tests, err := s.testRepository.FindTestsByTaskID(taskID)
	if err != nil {
		return TestsResult{}, fmt.Errorf("In TestService(RunTestOnFile'): %w", err)
	}

	if len(tests) == 0 {
		return TestsResult{}, TestsNotFoundError
	}

	var points int
	for _, test := range tests {
		timeout := timeLimit
		maxMemoryKB := memoryLimitInKB
		output, err := runCompiledCodeWithInput(fileName, test.Input, timeout, maxMemoryKB)
		if err != nil {
			if errors.Is(err, TimeLimitError) {
				return TestsResult{
					ResultCode: TimeLimitCode,
					Points:     points,
				}, nil
			}
			if errors.Is(err, ProgramError) {
				return TestsResult{
					ResultCode:  RuntimeErrorCode,
					Description: fmt.Sprintf("Error Info: %s Output: %s", err.Error(), output),
					Points:      points,
				}, nil
			}
			return TestsResult{}, fmt.Errorf("In TestService(RunTests): %w", err)
		}
		if output != test.ExpectedResult {
			return TestsResult{
				ResultCode:  IncorrectAnswerCode,
				Description: fmt.Sprintf("Test Failed: %d", test.ID),
				Points:      points,
			}, nil
		}
		points += test.Points
	}

	return TestsResult{
		ResultCode: SuccesCode,
		Points:     points,
	}, err
}

func runCompiledCodeWithInput(fileName string, input string, timeout time.Duration, maxMemoryKB int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "./"+fileName)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("In TestService(runCompiledCodeWithInput): %w", err)
	}
	defer stdin.Close()

	fmt.Fprintln(stdin, input)

	output, err := cmd.CombinedOutput()
	outputString := *(*string)(unsafe.Pointer(&output))
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok && status.Signaled() && status.Signal() == syscall.SIGKILL {
			return "", TimeLimitError
		}

		return outputString, fmt.Errorf("%w: %s", ProgramError, err)
	}

	return strings.ReplaceAll(outputString, "\n", ""), nil
}

func (s *TestService) GetTest(id int) (Test, error) {
	test, err := s.testRepository.FindTestByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return Test{}, ErrNotFound
		}
		return Test{}, fmt.Errorf("In TestService(GetTest): %w", err)
	}
	return test, nil
}

func (s *TestService) AddTest(taskID int, input string, expectedResult string, points int) error {
	err := s.testRepository.AddTest(taskID, input, expectedResult, points)
	if err != nil {
		return fmt.Errorf("In TestService(AddTest): %w", err)
	}
	return nil
}

func (s *TestService) DeleteTest(id int) error {
	err := s.testRepository.DeleteTest(id)
	if err != nil {
		return fmt.Errorf("In TestService(DeleteTest): %w", err)
	}
	return nil
}

func (s *TestService) UpdateTest(id int, newTest Test) error {
	err := s.testRepository.UpdateTest(id, newTest)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("In TestService(UpdateTest): %w", err)
	}
	return nil
}

func (s *TestService) GetTests() ([]Test, error) {
	tests, err := s.testRepository.GetTests()
	if err != nil {
		return nil, fmt.Errorf("In TestService(GetTests): %w", err)
	}
	return tests, nil
}

func (s *TestService) GetTestsByTaskID(taskID int) ([]Test, error) {
	tests, err := s.testRepository.FindTestsByTaskID(taskID)
	if err != nil {
		return nil, fmt.Errorf("In TestService(GetTestsByTaskID): %w", err)
	}
	return tests, nil
}
