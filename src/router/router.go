package router

import (
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/auth"
	"github.com/Taihenc/GO-Real-time-Leaderboard/src/handler"
)

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.ServePublic)
	mux.HandleFunc("/leaderboard", handler.GetScoreboard)
	mux.HandleFunc("/submit-score", handler.AddScore)
	mux.HandleFunc("/auth/register", auth.RegisterUser)
	mux.HandleFunc("/auth/login", auth.LoginUser)
}
