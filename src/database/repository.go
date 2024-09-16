package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetScoreboard() []string {
	_client := getRedisClient()

	val, err := _client.ZRevRange(ctx, "scoreboard", 0, 9).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func AddScore(game string, score int, player_name string) {
	_client := getRedisClient()

	_client.ZAdd(ctx, game, redis.Z{Score: float64(score), Member: player_name})
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
