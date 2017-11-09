package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/midorigreen/gprof/prof"
)

const port = ":8080"

func run() error {
	h := handler.New(
		&handler.Config{
			Schema:   &prof.Schema,
			Pretty:   true,
			GraphiQL: true,
		})

	http.Handle("/gq", h)
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "OK")
	})
	log.Printf("server run %s port", port)
	return http.ListenAndServe(port, nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
