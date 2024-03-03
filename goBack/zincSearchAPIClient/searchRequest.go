package zincSearchAPIClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	endpoint = "http://localhost:4080/api/Emails/_search"
	pageSize = 50 // Number of results per page

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
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	from := (page - 1) * pageSize
	query := createSearchQuery(searchTerm, from)

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.SetBasicAuth(creds.User, creds.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	
	fmt.Println("Response Status:", resp.Status)

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error parsing response: hits.hits field is not a slice of interfaces")
	}

	return convertHitsToMaps(hits), nil
}

func convertHitsToMaps(hits []interface{}) []map[string]interface{} {
	var values []map[string]interface{}
	for _, val := range hits {
		entry := val.(map[string]interface{})
		values = append(values, entry)
	}
	return values
}

func createSearchQuery(searchTerm string, from int) SearchQuery {
	if searchTerm == "" {
		return SearchQuery{
			SearchType: "matchall",
			From:       from,
			MaxResults: pageSize,
			Source:     []string{},
		}
	} else {
		return SearchQuery{
			SearchType: "match",
			Query: Query{
				Term: searchTerm,
			},
			From:       from,
			MaxResults: pageSize,
			Source:     []string{},
		}
	}
}
