package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/database"
	"github.com/Taihenc/GO-Real-time-Leaderboard/src/model"
)

var fileServer = http.FileServer(http.Dir("public"))

func ServePublic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./public/index.html")
		return
	}
	fileServer.ServeHTTP(w, r)
}

func AddScore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var player model.AddScoreRequest
		err := json.NewDecoder(r.Body).Decode(&player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		database.AddScore(player.Game, player.Score, player.PlayerName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Score added"))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
