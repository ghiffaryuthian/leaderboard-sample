package leaderboard

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type redisRepo struct {
	Redis    *redis.Client
	Name     string
	PageSize int
}

// NewRedisRepo implements the leaderboard interface using redis repository
func NewRedisRepo(redis *redis.Client, name string, pageSize int) Repository {
	return &redisRepo{
		Redis:    redis,
		Name:     name,
		PageSize: pageSize,
	}
}

// InsertUserScore will store a new member score into the leaderboard
func (r *redisRepo) InsertUserScore(ctx context.Context, username string, score float64) (*User, error) {
	var user User

	_, err := r.Redis.ZAdd(ctx, r.Name, &redis.Z{Score: score, Member: username}).Result()

	return &user, err
}

// GetUserScore will fetch member's details given username
func (r *redisRepo) GetUserRank(ctx context.Context, username string) (float64, error) {
	rankscore, err := r.Redis.ZRevRank(ctx, r.Name, username).Result()

	return float64(rankscore), err
}

func (r *redisRepo) TotalMembers(ctx context.Context) (int, error) {
	return 0, nil
}
