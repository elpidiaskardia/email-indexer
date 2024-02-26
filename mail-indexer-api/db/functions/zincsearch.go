package functions

import (
	"Api_Go/models"
	"Api_Go/env"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

)


func SendBulkToZincSearch(records []models.EmailData) error {
	bulkData := models.BulkData{
		Index:   "emails",
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		return err
	}

	req, err := makeRequest("POST", env.BaseURL+"_bulkv2", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return err
	}

	return handleResponse(resp, nil)
}

func CreateIndexerFromJsonFile(filepath string) (models.IndexerData, error) {
	var indexerData models.IndexerData

	file, err := os.Open(filepath)
	if err != nil {
		return indexerData, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexerData)
	if err != nil {
		return indexerData, err
	}

	return indexerData, nil
}

func CreateIndexOnZincSearch(indexerData models.IndexerData) error {
	jsonData, err := json.Marshal(indexerData)
	if err != nil {
		return err
	}

	req, err := makeRequest("POST", env.BaseURL+"index", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create indexer, status code: %d", resp.StatusCode)
	}

	return nil
}

func DeleteIndexOnZincSearch(indexName string) error {
	req, err := makeRequest("DELETE", env.BaseURL+"index/"+indexName, nil)
	if err != nil {
		return err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete indexer, status code: %d", resp.StatusCode)
	}

	log.Println("Index deleted successfully")
	return nil
}


