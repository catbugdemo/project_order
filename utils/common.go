package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// HttpGet
func HttpGet(url string, values interface{}) error {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	if err = json.NewDecoder(resp.Body).Decode(values); err != nil {
		return err
	}
	return nil
}
