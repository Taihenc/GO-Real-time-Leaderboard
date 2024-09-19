module github.com/Taihenc/GO-Real-time-Leaderboard

go 1.21

require (
	github.com/Taihenc/GO-Real-time-Leaderboard/src/auth v0.0.0
	github.com/Taihenc/GO-Real-time-Leaderboard/src/database v0.0.0
	github.com/Taihenc/GO-Real-time-Leaderboard/src/handler v0.0.0
	github.com/Taihenc/GO-Real-time-Leaderboard/src/model v0.0.0
	github.com/Taihenc/GO-Real-time-Leaderboard/src/multiplexer v0.0.0
	github.com/Taihenc/GO-Real-time-Leaderboard/src/router v0.0.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.6.1
	golang.org/x/crypto v0.27.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/database => ./src/database

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/auth => ./src/auth

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/handler => ./src/handler

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/model => ./src/model

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/multiplexer => ./src/multiplexer

replace github.com/Taihenc/GO-Real-time-Leaderboard/src/router => ./src/router
