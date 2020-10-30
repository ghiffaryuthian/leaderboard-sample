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
func (r *redisRepo) InsertUserScore(ctx context.Context, username string, score float64) error {
	_, err := r.Redis.ZAdd(ctx, r.Name, &redis.Z{Score: score, Member: username}).Result()

	return err
}

// GetUserRank will fetch member's details given username
func (r *redisRepo) GetUserRank(ctx context.Context, username string) (int64, error) {
	rank, err := r.Redis.ZRevRank(ctx, r.Name, username).Result()
	// convert 0-base to 1-base
	rank++

	return rank, err
}

// GetUserScore will fetch member's details given username
func (r *redisRepo) GetUserScore(ctx context.Context, username string) (float64, error) {
	score, err := r.Redis.ZScore(ctx, r.Name, username).Result()

	return float64(score), err
}

func (r *redisRepo) TotalMembers(ctx context.Context) (int, error) {
	return 0, nil
}
