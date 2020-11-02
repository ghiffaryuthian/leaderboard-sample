package leaderboard

import (
	"context"
)

// Service is interface for leaderboard business logic
type Service interface {
	RankMember(ctx context.Context, username string, score float64) (*User, error)
	GetMemberDetails(ctx context.Context, username string) (*User, error)
}

type service struct {
	repo Repository
}

// NewService returns interfaces for services implementation
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// RankMember inserts user score if it's higher and return final user rank
func (s *service) RankMember(ctx context.Context, username string, score float64) (*User, error) {
	scoreOld, err := s.repo.GetUserScore(ctx, username)
	if err != nil {
		return nil, err
	}
	if score > scoreOld {
		if err := s.repo.InsertUserScore(ctx, username, score); err != nil {
			return nil, err
		}
	} else {
		score = scoreOld
	}
	rank, err := s.repo.GetUserRank(ctx, username)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Rank:     rank,
		Score:    score,
	}
	return user, err
}

//GetMemberDetails returns member details
func (s *service) GetMemberDetails(ctx context.Context, username string) (*User, error) {
	rank, err := s.repo.GetUserRank(context.TODO(), username)
	if err != nil {
		return nil, err
	}
	score, err := s.repo.GetUserScore(context.TODO(), username)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username: username,
		Rank:     rank,
		Score:    score,
	}
	return user, err
}
