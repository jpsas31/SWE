package chiAPI

import (
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)

	fmt.Printf("Error: %v\n", err)
}
