package zincSearchAPIClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	ZincHost     = "http://localhost:4080"
	BulkEndpoint = "/api/_bulk"
)

// BulkIndex sends the provided email entries to the Zinc API for bulk indexing.
// It loads credentials from the secret file, constructs the necessary HTTP request with appropriate headers,
// and sends the request using the zincSearch _bulk api endpoint. It returns an error if any operation fails.
func BulkIndex(entries []byte) error {
	creds, err := LoadCredentials(SecretFilePath)
	if err != nil {
		return fmt.Errorf("failed to load credentials: %w", err)
	}

	zincURL := ZincHost + BulkEndpoint

	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer(entries))
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
