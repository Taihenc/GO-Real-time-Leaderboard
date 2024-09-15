package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/multiplexer"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	multiplexer.Initialize()

	fmt.Println("Server is running on port", PORT, "at http://localhost:"+PORT)
	if err := http.ListenAndServe(":"+PORT, multiplexer.Mux); err != nil {
		fmt.Println(err)
	}
}
