package leaderboard

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

type redisRepo struct {
	Redis    *redis.Client
	Username string
	PageSize int
}

// NewRedisRepo implements the leaderboard interface using redis repository
func NewRedisRepo(redis *redis.Client, name string, pageSize int) Repository {
	return &redisRepo{
		Redis:    redis,
		Username: name,
		PageSize: pageSize,
	}
}

// InsertUserScore stores a new member score into the leaderboard
func (r *redisRepo) InsertUserScore(ctx context.Context, username string, score float64) error {
	_, err := r.Redis.ZAdd(ctx, r.Username, &redis.Z{Score: score, Member: username}).Result()

	return err
}

// GetUserRank fetchs member's rank given username
func (r *redisRepo) GetUserRank(ctx context.Context, username string) (int64, error) {
	rank, err := r.Redis.ZRevRank(ctx, r.Username, username).Result()
	// convert 0-base to 1-base
	rank++

	return rank, err
}

// GetUserScore fetchs member's score given username
func (r *redisRepo) GetUserScore(ctx context.Context, username string) (float64, error) {
	score, err := r.Redis.ZScore(ctx, r.Username, username).Result()
	if err == redis.Nil {
		return score, nil
	}

	return score, err
}

// TotalMembers give the count of all unique user on leaderboard
func (r *redisRepo) TotalMembers(ctx context.Context) (int64, error) {
	memberCount, err := r.Redis.ZCount(ctx, r.Username, "-inf", "+inf").Result()

	return memberCount, err
}
