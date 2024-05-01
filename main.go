package main

import (
	"net/http"
	"the-word/handlers"
)

func registerHandlers() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.Handle("/static/", handlers.StaticHandler())
}

func main() {
	registerHandlers()
	http.ListenAndServe(":9090", nil)
}
