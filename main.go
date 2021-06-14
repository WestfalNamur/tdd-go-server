package main

import (
	"log"
	"net/http"
	"tdd-go-server/handler"
)

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	// server is a reference to PlayerServer and takes a reference to
	// InMemoryPlayerStore.
	server := &handler.PlayerServer{Store: &InMemoryPlayerStore{}}

	// ListenAndServe listens on the TCP network address addr and then calls
	// Serve with handler to handle requests on incoming connections. Accepted
	// connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":5000", server))
}
