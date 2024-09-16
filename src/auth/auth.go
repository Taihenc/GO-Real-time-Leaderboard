package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/database"
	"github.com/Taihenc/GO-Real-time-Leaderboard/src/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		// Store the hashed password in Redis
		err = database.RegisterUser(user.Username, hashedPassword)

		if err != nil {
			if err.Error() == "User already exists" {
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered successfully!"))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
