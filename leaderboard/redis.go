package leaderboard

import (
	redis "github.com/go-redis/redis/v8"
)

// NewRedisRepo implements the leaderboard interface using redis repository
func NewRedisRepo(redis *redis.Client) Repository {

}
