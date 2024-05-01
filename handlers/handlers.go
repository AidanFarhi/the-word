package handlers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func StaticHandler() http.Handler {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	return handler
}
