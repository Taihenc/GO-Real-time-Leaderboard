package database

import (
	"fmt"
	"time"

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

	// set last update time
	err = _client.Set(ctx, "lastUpdateTime:"+record.Game, time.Now(), 0).Err()
	if err != nil {
		fmt.Println("Error setting last update time")
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

func GetGameList() ([]string, error) {
	_client := getRedisClient()

	// get all games name from LIST with key "GameList"
	games, err := _client.LRange(ctx, "GameList", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return games, nil
}

func GetLastUpdateTime(game string) (string, error) {
	_client := getRedisClient()

	lastUpdateTime, err := _client.Get(ctx, "lastUpdateTime:"+game).Result()
	if err != nil {
		fmt.Println("Error getting last update time")
		return "", err
	}
	return lastUpdateTime, nil
}
