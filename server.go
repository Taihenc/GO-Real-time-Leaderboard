package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

var fileServer = http.FileServer(http.Dir("public"))

func servePublic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./public/index.html")
		return
	}
	fileServer.ServeHTTP(w, r)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", servePublic)

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux); err != nil {
		fmt.Println(err)
	}
}
