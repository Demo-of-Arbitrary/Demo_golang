package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		response.WriteHeader(http.StatusAccepted)
		return
	}
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}
