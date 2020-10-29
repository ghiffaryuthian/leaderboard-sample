package leaderboard

import (
	"context"
)

// Repository holds interface for leaderboard functions
type Repository interface {
	RankMember(ctx context.Context, username string, score int) (User, error)
	GetMember(ctx context.Context, username string) (User, error)
	TotalMembers(ctx context.Context) (int, error)
}