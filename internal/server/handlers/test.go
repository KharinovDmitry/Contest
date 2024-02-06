package handlers

import (
	. "contest/internal/domain"
	"contest/internal/services/testCrudService"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

type ITestCrudService interface {
	GetTest(id int) (Test, error)
	AddTest(taskID int, input string, expectedResult string, points int) error
	DeleteTest(id int) error
	UpdateTest(id int, newTest Test) error
	GetTests() ([]Test, error)
	GetTestsByTaskID(taskID int) ([]Test, error)
}

type getTestsResponse struct {
	Tests []testDTO `json:"tests"`
}

type addTestRequest struct {
	TaskID         int    `json:"taskID,string"`
	Input          string `json:"input"`
	ExpectedResult string `json:"expectedResult"`
	Points         int    `json:"points,string"`
}

type testDTO struct {
	ID             int    `json:"id"`
	TaskID         int    `json:"taskID"`
	Input          string `json:"input"`
	ExpectedResult string `json:"expectedResult"`
	Points         int    `json:"points"`
}

func testsToTestsDTO(tests []Test) []testDTO {
	res := make([]testDTO, len(tests))
	for i, test := range tests {
		res[i] = testDTO(test)
	}
	return res
}

func AddTest(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var test addTestRequest
		if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Error(err.Error())
			return
		}

		if err := testService.AddTest(test.TaskID, test.Input, test.ExpectedResult, test.Points); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTest(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = testService.DeleteTest(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateTest(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var testDTO testDTO
		if err := json.NewDecoder(r.Body).Decode(&testDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = testService.UpdateTest(id, Test(testDTO))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetTest(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		test, err := testService.GetTest(id)
		if err != nil {
			log.Error(err.Error())

			if errors.Is(err, testCrudService.ErrNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		response, err := json.Marshal(test)
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func GetTests(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		tests, err := testService.GetTests()
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(getTestsResponse{
			Tests: testsToTestsDTO(tests),
		})
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func GetTestsByTaskID(testService ITestCrudService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		taskID, err := strconv.Atoi(mux.Vars(r)["task_id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tests, err := testService.GetTestsByTaskID(taskID)
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(getTestsResponse{
			Tests: testsToTestsDTO(tests),
		})
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
