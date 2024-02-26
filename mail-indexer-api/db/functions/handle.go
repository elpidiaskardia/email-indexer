package functions

import (
	"Api_Go/env"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(env.AuthUsername, env.AuthPassword)
	req.Header.Set("Content-Type", env.ContentType)
	req.Header.Set("User-Agent", env.UserAgent)
	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func handleResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}



