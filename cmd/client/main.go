package main

import (
	"github.com/Anhbman/microservice-server-cake/rpc/query"
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := query.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &query.Size{Inches: 12})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("I have a nice new hat: %+v", hat)
}
