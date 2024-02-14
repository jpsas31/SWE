package zincSearchAPIClient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	ENDPOINT  = "http://localhost:4080/api/Emails/_search"
	PAGE_SIZE = 50 // Number of results per page

)

type SearchQuery struct {
	SearchType string   `json:"search_type"`
	Query      Query    `json:"query"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	Source     []string `json:"_source"`
}

type Query struct {
	Term string `json:"term"`
}

func Search(page int, searchTerm string) ([]map[string]interface{}, error) {
	creds, err := LoadCredentials(SecretFilePath)
	if err != nil {
		log.Printf("Error loading credentials: %v", err)
		return nil, err
	}

	from := (page - 1) * PAGE_SIZE
	var query SearchQuery
	if searchTerm == "" {
		query = SearchQuery{
			SearchType: "matchall",
			From:       from,
			MaxResults: PAGE_SIZE,
			Source:     []string{},
		}
	} else {
		query = SearchQuery{
			SearchType: "match",
			Query: Query{
				Term: searchTerm,
			},
			From:       from,
			MaxResults: PAGE_SIZE,
			Source:     []string{},
		}
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		log.Printf("Error marshalling query: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", ENDPOINT, strings.NewReader(string(jsonData)))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.SetBasicAuth(creds.User, creds.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error executing request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return nil, err
	}

	fmt.Println("Response Status:", resp.Status)

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		err := fmt.Errorf("hits.hits field is not a slice of interfaces")
		log.Printf("Error parsing response: %v", err)
		return nil, err
	}

	var values = []map[string]interface{}{}
	for _, val := range hits {
		entry := val.(map[string]interface{})
		values = append(values, entry)
	}

	return values, nil
}
