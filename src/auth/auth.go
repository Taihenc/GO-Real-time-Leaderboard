package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/database"
	"github.com/Taihenc/GO-Real-time-Leaderboard/src/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

// JWT secret key (keep this secret)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.UserPassword
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
		// Generate a JWT token if registration is successful

		tokenString, err := GenerateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful!",
			"token":   tokenString,
		})
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.UserPassword
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Fetch hashed password from Redis
		storedHash, err := database.GetHashedPassword(user.Username)
		if err == redis.Nil {
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, "Error fetching user", http.StatusInternalServerError)
			return
		}

		// Compare the hashed password with the password provided by the user
		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate a JWT token if the credentials are correct
		tokenString, err := GenerateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful!",
			"token":   tokenString,
		})
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
