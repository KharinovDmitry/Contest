package runTestService_test

import (
	"contest/internal/domain"
	"contest/internal/services/runTestService"
	mock_executorFactory "contest/internal/services/runTestService/mocks"
	mock_storage "contest/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initMockRepo(cntrl *gomock.Controller) *mock_storage.MockTestRepository {
	repo := mock_storage.NewMockTestRepository(cntrl)
	repo.EXPECT().FindTestsByTaskID(1).Return(
		[]domain.Test{
			{
				ID:             1,
				TaskID:         1,
				Input:          "1",
				ExpectedResult: "1",
				Points:         1,
			},
			{
				ID:             2,
				TaskID:         1,
				Input:          "2",
				ExpectedResult: "4",
				Points:         1,
			},
			{
				ID:             3,
				TaskID:         1,
				Input:          "3",
				ExpectedResult: "9",
				Points:         1,
			},
		}, nil)

	return repo
}

func initMockExecutorFactory(cntrl *gomock.Controller) *mock_executorFactory.MockIExecutorFactory {
	mockFactory := mock_executorFactory.NewMockIExecutorFactory(cntrl)
	mockExecutor := mock_executorFactory.NewMockIExecutor(cntrl)
	mockExecutor.EXPECT().Execute("1", gomock.Any(), gomock.Any()).Return("1", nil)
	mockExecutor.EXPECT().Execute("2", gomock.Any(), gomock.Any()).Return("4", nil)
	mockExecutor.EXPECT().Execute("3", gomock.Any(), gomock.Any()).Return("9", nil)

	mockExecutor.EXPECT().Close()

	mockFactory.EXPECT().NewExecutor(gomock.Any(), gomock.Any()).Return(mockExecutor, nil)

	return mockFactory
}

func Test(t *testing.T) {
	cntrl := gomock.NewController(t)
	testRepo := initMockRepo(cntrl)
	executorFactory := initMockExecutorFactory(cntrl)
	service := runTestService.NewRunTestService(executorFactory, testRepo)

	actual, err := service.RunTest(1, domain.CPP, "", 1024, 100)
	assert.Nil(t, err)
	expected := domain.TestsResult{
		ResultCode:  "SC",
		Description: "",
		Points:      3,
	}
	assert.Equal(t, expected, actual)
}
