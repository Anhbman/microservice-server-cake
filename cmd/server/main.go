package main

import (
	"cake/internal/haberdasherserver"
	"cake/rpc/query"
	"net/http"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := query.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
