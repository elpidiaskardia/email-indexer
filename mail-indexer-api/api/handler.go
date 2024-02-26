package api

import (
	"Api_Go/models"
	"Api_Go/services"
	"encoding/json"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request, getResultFunc func(models.PaginationParams) (models.ApiResult, error)) {
	var pagination models.PaginationParams
	err := json.NewDecoder(r.Body).Decode(&pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := getResultFunc(pagination)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}

	status := http.StatusOK
	message := "Data Found"
	if result.Result.Total == 0 {
		status = http.StatusNotFound
		message = "No Data Found"
	}

	result.Status = "Success"
	result.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func GetEmailByParameter(w http.ResponseWriter, r *http.Request) {
	searchParameter := r.URL.Query().Get("searchParameter")
	getResultFunc := func(pagination models.PaginationParams) (models.ApiResult, error) {
		return services.GetEmailByParameter(searchParameter, pagination)
	}
	handleRequest(w, r, getResultFunc)
}

func GetAllEmails(w http.ResponseWriter, r *http.Request) {
	getResultFunc := func(pagination models.PaginationParams) (models.ApiResult, error) {
		return services.GetAllEmails(pagination)
	}
	handleRequest(w, r, getResultFunc)
}
