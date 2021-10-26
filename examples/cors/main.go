package main

import (
	"log"
	"net/http"

	"github.com/clarify/rested/resource"
	"github.com/clarify/rested/rest"
	"github.com/rs/cors"
)

func main() {
	index := resource.NewIndex()

	// configure your resources

	api, err := rest.NewHandler(index)
	if err != nil {
		log.Fatalf("Invalid API configuration: %s", err)
	}

	handler := cors.Default().Handler(api)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
