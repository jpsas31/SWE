package zincSearchAPIClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	
	docEndpoint = "/api/_doc"
)

func DocIndex(entry []byte) error {
	creds, err := LoadCredentials(SecretFilePath)
	if err != nil {
		return fmt.Errorf("failed to load credentials: %w", err)
	}

	zincURL := zincHost + docEndpoint

	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer(entry))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.SetBasicAuth(creds.User, creds.Password)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))

	return nil
}
