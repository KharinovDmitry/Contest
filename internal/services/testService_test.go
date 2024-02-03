package services

import (
	"contest/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSucces(t *testing.T) {
	compiler := mocks.MockCompiler{}
	testRepo := mocks.MockTestRepository{}
	testService := NewTestService(compiler, testRepo)

	actual, err := testService.RunTestOnFile("testFiles/testSucces", 1, false)
	expected := domain.TestsResult{
		ResultCode:  domain.SuccesCode,
		Description: "",
		Points:      3,
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestTimeLimit(t *testing.T) {
	compiler := mocks.MockCompiler{}
	testRepo := mocks.MockTestRepository{}
	testService := NewTestService(compiler, testRepo)

	actual, err := testService.RunTestOnFile("testFiles/testTimeLimit", 1, false)
	expected := domain.TestsResult{
		ResultCode:  domain.TimeLimitCode,
		Description: "",
		Points:      0,
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestNotFullSolution(t *testing.T) {
	compiler := mocks.MockCompiler{}
	testRepo := mocks.MockTestRepository{}
	testService := NewTestService(compiler, testRepo)

	actual, err := testService.RunTestOnFile("testFiles/testNotFullSolution", 1, false)
	expected := domain.TestsResult{
		ResultCode:  domain.IncorrectAnswerCode,
		Description: "Test Failed: 2",
		Points:      1,
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestRuntimeError(t *testing.T) {
	compiler := mocks.MockCompiler{}
	testRepo := mocks.MockTestRepository{}
	testService := NewTestService(compiler, testRepo)

	actual, err := testService.RunTestOnFile("testFiles/testRuntimeError", 1, false)
	expected := domain.TestsResult{
		ResultCode:  domain.RuntimeErrorCode,
		Description: "Error Info: Program error: signal: segmentation fault (core dumped) Output: ",
		Points:      0,
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
