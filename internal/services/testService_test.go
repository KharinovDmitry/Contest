package services

import (
	"contest/internal/compiler/mocks"
	"contest/internal/domain"
	mock_storage "contest/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InitMockRepo(cntrl *gomock.Controller) *mock_storage.MockTestRepository {
	testRepo := mock_storage.NewMockTestRepository(cntrl)

	testRepo.EXPECT().FindTestsByTaskID(1).Return([]domain.Test{
		domain.Test{
			ID:             1,
			TaskID:         1,
			Input:          "1",
			ExpectedResult: "1",
			Points:         1,
		},
		domain.Test{
			ID:             2,
			TaskID:         1,
			Input:          "2",
			ExpectedResult: "4",
			Points:         1,
		},
		domain.Test{
			ID:             2,
			TaskID:         1,
			Input:          "3",
			ExpectedResult: "9",
			Points:         1,
		},
	}, nil)

	return testRepo
}

func TestSucces(t *testing.T) {
	cntrl := gomock.NewController(t)
	compiler := mock_compiler.NewMockCompiler(cntrl)
	testRepo := InitMockRepo(cntrl)
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
	cntrl := gomock.NewController(t)
	compiler := mock_compiler.NewMockCompiler(cntrl)
	testRepo := InitMockRepo(cntrl)
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
	cntrl := gomock.NewController(t)
	compiler := mock_compiler.NewMockCompiler(cntrl)
	testRepo := InitMockRepo(cntrl)
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
	cntrl := gomock.NewController(t)
	compiler := mock_compiler.NewMockCompiler(cntrl)
	testRepo := InitMockRepo(cntrl)
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
