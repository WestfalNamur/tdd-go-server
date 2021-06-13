package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

// Implement ServerHTTP method on PlayerServer.
// A Handler responds to an HTTP request. ServeHTTP should write reply headers
// and data to the ResponseWriter and then return. Returning signals that the
// request is finished; it is not valid to use the ResponseWriter or read from
// the Request.Body after or concurrently with the completion of the ServeHTTP
// call.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// return string without the provided leading prefix string.
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	// Store and methode injected during test.
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
