package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer is
func PlayerServer(response http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	fmt.Fprintf(response, GetPlayerScore(player))
}

// GetPlayerScore is
func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}
	return ""
}
