package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

const jsonContentType = "application/json"

func (p *PlayerServer) leagueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", jsonContentType)
	json.NewEncoder(rw).Encode(p.store.GetLeague())
	rw.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(rw, player)
	case http.MethodGet:
		p.showScore(rw, player)
	}
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(rw, score)
}

func (p *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	rw.WriteHeader(http.StatusAccepted)
	p.store.RecordWin(player)
}
