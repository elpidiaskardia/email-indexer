package models

type PaginationParams struct {
    CurrentPage int `json:"current_page"`
    PageSize    int `json:"page_size"`
    TotalRecords int `json:"total_records"`
}