package models

type ApiResult struct {
	Status  string  `json:"status"`
	Result  EmailResult  `json:"result"`
	Message  string  `json:"message"`
}

type EmailResult struct {
	Total  int `json:"total"`
	Data []EmailData  `json:"data"`
}