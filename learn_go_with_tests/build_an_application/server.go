package server

import (
	"fmt"
	"net/http"
)

// PlayerServer is
func PlayerServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "20")
}
