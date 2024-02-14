package zincSearchAPIClient

import (
	"encoding/json"
	"os"
)

// Credentials struct to hold user and password
type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

const (
	SecretFilePath = "secret.json"
)

// LoadCredentials reads credentials from a file
func LoadCredentials(filepath string) (*Credentials, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	creds := &Credentials{}
	err = decoder.Decode(creds)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
