package database

import (
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
