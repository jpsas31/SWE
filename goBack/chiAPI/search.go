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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var decoded searchRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Printf("Requested %s\n", decoded.SearchTerm)

	result, err := zincSearchAPIClient.Search(decoded.Page, decoded.SearchTerm)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
	response, err := json.Marshal(result)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}

	_, err = w.Write(response)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
