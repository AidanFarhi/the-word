package handlers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func StaticHandler() http.Handler {
	staticBaseDir := "/static/"
	fullPath := "static"
	handler := http.StripPrefix(staticBaseDir, http.FileServer(http.Dir(fullPath)))
	return handler
}
