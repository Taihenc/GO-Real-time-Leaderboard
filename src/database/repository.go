package database

import (
	"fmt"

	"github.com/Taihenc/GO-Real-time-Leaderboard/src/model"
	"github.com/redis/go-redis/v9"
)

func GetScoreboard(game string) ([]model.LeaderboardRecord, error) {
	_client := getRedisClient()

	records, err := _client.ZRevRangeWithScores(ctx, game, 0, 9).Result()
	if err != nil {
		fmt.Println("Error getting leaderboard")
		return nil, err
	}

	var leaderboard []model.LeaderboardRecord
	for _, record := range records {
		leaderboard = append(leaderboard, model.LeaderboardRecord{
			Game:       game,
			PlayerName: record.Member.(string),
			Score:      int(record.Score),
		})
	}
	return leaderboard, nil
}

func AddScore(record model.LeaderboardRecord) error {
	_client := getRedisClient()

	err := _client.ZAdd(ctx, record.Game, redis.Z{Score: float64(record.Score), Member: record.PlayerName}).Err()
	if err != nil {
		fmt.Println("Error adding score")
		return err
	}
	fmt.Println("Score added successfully!")
	return nil
}

func RegisterUser(username string, hashedPassword []byte) error {
	_client := getRedisClient()

	// check if user already exists
	_, err := _client.Get(ctx, "user:"+username).Result()
	if err == nil {
		fmt.Println("User already exists")
		return fmt.Errorf("User already exists")
	}
	err = _client.Set(ctx, "user:"+username, hashedPassword, 0).Err()
	if err != nil {
		fmt.Println("Error saving user")
		return err
	}

	fmt.Println("User registered successfully!")
	return nil
}

func GetHashedPassword(username string) ([]byte, error) {
	_client := getRedisClient()

	hashedPassword, err := _client.Get(ctx, "user:"+username).Bytes()
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
