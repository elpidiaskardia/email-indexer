package api

import (
	"Api_Go/models"
	"Api_Go/services"
	"encoding/json"
	"net/http"
)

// Responds with the result of the function using the result model.
func handleRequest(writer http.ResponseWriter, request *http.Request, getResultFunc func(models.PaginationParams) (models.ApiResult, error)) {
	var pagination models.PaginationParams
	if err := json.NewDecoder(request.Body).Decode(&pagination); err != nil {
		handleError(writer, http.StatusBadRequest, "fail", err.Error())
		return
	}

	result, err := getResultFunc(pagination)
	if err != nil {
		handleError(writer, http.StatusInternalServerError, "fail", err.Error())
		return
	}

	status := http.StatusOK
	message := "Data Found"
	result.Status = "Success"
	result.Message = message

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(result)

}

// configures the response in case of an error.
func handleError(writer http.ResponseWriter, status int, resultStatus, message string) {
	var result models.ApiResult
	result.Status = resultStatus
	result.Message = message

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(result)
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
