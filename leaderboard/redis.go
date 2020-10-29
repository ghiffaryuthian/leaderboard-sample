package leaderboard

import (
	"context"
	"errors"
	"fmt"

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

// RankMember will store a new member score into the leaderboard
func (r *redisRepo) RankMember(ctx context.Context, username string, score float64) (*User, error) {
	var user User
	if val := r.Redis.ZAdd(ctx, r.Name, &redis.Z{Score: score, Member: username}); val == nil {
		msg := fmt.Sprintf("Failed to insert member %s with score %f to %s", username, score, r.Name)
		err := errors.New(msg)
		return &user, err
	}
	return &user, nil
}

// GetMember will fetch member's details given username
func (r *redisRepo) GetMember(ctx context.Context, username string) (*User, error) {
	var user User
	return &user, nil
}

func (r *redisRepo) TotalMembers(ctx context.Context) (int, error) {
	return 0, nil
}
