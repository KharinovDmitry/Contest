package handlers

import (
	. "contest/internal/domain"
	"contest/internal/services"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type runTestRequest struct {
	TaskID          int      `json:"task_id,string"`
	Language        Language `json:"language"`
	Code            string   `json:"code"`
	MemoryLimitInKb int      `json:"memoryLimitInMs,string"`
	TimeLimitInMs   int      `json:"timeLimitInKb,string"`
}

type runTestResponse struct {
	ResultCode  string `json:"result_code"`
	Description string `json:"description"`
	Points      int    `json:"points,string"`
}

func RunTest(testService services.ITestService, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Info("RunTest request received")

		var request runTestRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := testService.RunTest(RunTestRequest(request))
		if err != nil {
			if errors.Is(err, services.UnknownLanguageError) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(runTestResponse{
			ResultCode:  string(result.ResultCode),
			Description: result.Description,
			Points:      result.Points,
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
