package models

type BulkData struct {
	Index   string      `json:"index"`
	Records []EmailData `json:"records"`
}