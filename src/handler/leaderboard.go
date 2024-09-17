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

func GetScoreboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		gameName := r.URL.Query().Get("game")
		if gameName == "" {
			http.Error(w, "Game name is required", http.StatusBadRequest)
			return
		}

		var scoreboard = []model.LeaderboardRecord{}
		scoreboard, err := database.GetScoreboard(gameName)
		if err != nil {
			http.Error(w, "Error getting leaderboard", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(scoreboard)
		return
	}
}

func GetGameList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var games = []string{}
		games, err := database.GetGameList()
		if err != nil {
			http.Error(w, "Error getting game list", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(games)
		return
	}
}

func GetLastUpdateTime(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		gameName := r.URL.Query().Get("game")
		if gameName == "" {
			http.Error(w, "Game name is required", http.StatusBadRequest)
			return
		}

		lastUpdateTime, err := database.GetLastUpdateTime(gameName)
		if err != nil {
			http.Error(w, "Error getting last update time", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lastUpdateTime)
		return
	}
}
