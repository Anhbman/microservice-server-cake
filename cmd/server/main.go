package main

import (
	"github.com/Anhbman/microservice-server-cake/internal/haberdasherserver"
	"github.com/Anhbman/microservice-server-cake/rpc/query"
	"net/http"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := query.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
