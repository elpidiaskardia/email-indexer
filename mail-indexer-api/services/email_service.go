package services

import (
	"fmt"
	"Api_Go/db/functions"
	"Api_Go/models"
)
//// HandleEmailsResponse handles the response of email search.
func HandleEmailsResponse(emails models.ZincsearchResult, err error) (models.ApiResult, error) {
	if err != nil {
		return models.ApiResult{}, fmt.Errorf("failed to retrieve emails: %w", err)
	}
	result := models.ApiResult{
		Result: models.EmailResult{
			Total: emails.Hits.Total.Value,
			Data:  make([]models.EmailData, len(emails.Hits.Hits)),
		},
	}
	for i, emailHit := range emails.Hits.Hits {
		result.Result.Data[i] = emailHit.Source
	}
	return result, nil
}
// GetEmailByParameter retrieves emails based on a search parameter and pagination.
func GetEmailByParameter(searchParameter string, pagination models.PaginationParams) (models.ApiResult, error) {
	emails, err := functions.GetListEmails(searchParameter, pagination)
	return HandleEmailsResponse(emails, err)
}
// GetAllEmails retrieves all emails with pagination.
func GetAllEmails(pagination models.PaginationParams) (models.ApiResult, error) {
	emails, err := functions.GetAllEmails(pagination)
	return HandleEmailsResponse(emails, err)
}
