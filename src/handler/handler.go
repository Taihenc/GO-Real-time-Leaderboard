package handler

import (
	"net/http"
)

var fileServer = http.FileServer(http.Dir("public"))

func ServePublic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./public/index.html")
		return
	}
	fileServer.ServeHTTP(w, r)
}
