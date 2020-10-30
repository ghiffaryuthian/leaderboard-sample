package leaderboard

import (
	"context"
)

// Repository holds interface for leaderboard functions
type Repository interface {
	InsertUserScore(ctx context.Context, username string, score float64) error
	GetUserRank(ctx context.Context, username string) (int64, error)
	GetUserScore(ctx context.Context, username string) (float64, error)
	TotalMembers(ctx context.Context) (int, error)
}
