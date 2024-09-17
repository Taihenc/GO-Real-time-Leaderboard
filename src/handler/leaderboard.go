package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/database"
	"github.com/Taihenc/GO-Real-time-Leaderboard/src/model"
)

func AddScore(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var player model.LeaderboardRecord
		err := json.NewDecoder(r.Body).Decode(&player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		err = database.AddScore(
			model.LeaderboardRecord{
				Game:       player.Game,
				PlayerName: player.PlayerName,
				Score:      player.Score,
			})
		if err != nil {
			http.Error(w, "Error adding score", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Score added"))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
