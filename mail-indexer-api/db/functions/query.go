package functions

import (
	"Api_Go/env"
	"Api_Go/models"
	"fmt"
	"strings"
)



func GetListEmails(searchParameter string, pagination models.PaginationParams) (models.ZincsearchResult, error) {
	query := fmt.Sprintf(`{
		"search_type": "matchphrase",
        "query": {
            "term": "%s"
        },
		"sort_fields": ["-@timestamp"],
		"from":  %d,
        "max_results": %d
    }`, searchParameter, pagination.CurrentPage, pagination.PageSize)

	req, err := makeRequest("POST", env.BaseURL+"emails/_search", strings.NewReader(query))
	if err != nil {
		return models.ZincsearchResult{}, err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return models.ZincsearchResult{}, err
	}

	var result models.ZincsearchResult
	if err := handleResponse(resp, &result); err != nil {
		return models.ZincsearchResult{}, err
	}

	return result, nil
}

func GetAllEmails(pagination models.PaginationParams) (models.ZincsearchResult, error) {
	query := fmt.Sprintf(`{
		"search_type": "matchall",
		"sort_fields": ["-@timestamp"],
		"from":  %d,
        "max_results": %d
    }`, (pagination.CurrentPage-1)*pagination.PageSize, pagination.PageSize)

	req, err := makeRequest("POST", env.BaseURL+"emails/_search", strings.NewReader(query))
	if err != nil {
		return models.ZincsearchResult{}, err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return models.ZincsearchResult{}, err
	}

	var result models.ZincsearchResult
	if err := handleResponse(resp, &result); err != nil {
		return models.ZincsearchResult{}, err
	}

	return result, nil
}
