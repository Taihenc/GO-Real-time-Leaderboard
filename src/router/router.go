package router

import (
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/handler"
)

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.ServePublic)
	mux.HandleFunc("/add_score", handler.AddScore)
}
