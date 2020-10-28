package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
}

type Player struct {
	Name string
	Wins int
}
type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
	GetLeague() League
}

const jsonContentType = "application/json"
const htmlTemplatePath = "game.html"

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template %s %v", htmlTemplatePath, err)
	}

	p.template = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p, nil
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	p.store.RecordWin(string(winnerMsg))
}
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWins(w, player)
	}
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) processWins(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(response http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}
