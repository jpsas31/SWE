package chiAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jpsas31/SWE/indexer/zincSearchAPIClient"
)

type searchRequest struct {
	SearchTerm string `json:"search_term"`
	Page       int    `json:"page"`
}
type searchResponse struct {
	Results  []map[string]interface{} `json:"results"`
	TotalPages   int      `json:"pages"`
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var decoded searchRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Printf("Requested %s\n", decoded.SearchTerm)

	result, pages, err := zincSearchAPIClient.Search(decoded.Page, decoded.SearchTerm)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
	response := searchResponse{
		Results: result,
		TotalPages: pages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}
