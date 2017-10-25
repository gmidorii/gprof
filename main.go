package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/midorigreen/gprof/prof"
)

func run() error {
	h := handler.New(
		&handler.Config{
			Schema:   &prof.Schema,
			Pretty:   true,
			GraphiQL: true,
		})

	http.Handle("/gq", h)
	return http.ListenAndServe(":8080", nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
