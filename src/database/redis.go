package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func getRedisClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),     // use default Addr
			Password: os.Getenv("REDIS_PASSWORD"), // no password
			DB:       0,                           // use default DB
		})
		// Test the connection
		_, err := client.Ping(ctx).Result()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
		}
	}
	return client
}
