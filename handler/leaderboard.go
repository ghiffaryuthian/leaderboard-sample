package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ghiffaryuthian/leaderboard-sample/leaderboard"
	"github.com/labstack/echo/v4"
)

type lbHandler struct {
	LeaderboardRepo leaderboard.Repository
}

// MakeLeaderboardHandler registers endpoints for healthz
func MakeLeaderboardHandler(e *echo.Echo, lbRepo leaderboard.Repository) {
	h := lbHandler{LeaderboardRepo: lbRepo}
	e.GET("/leaderboards/ranks/:username", h.getUserDetails)
}

func (h *lbHandler) getUserDetails(c echo.Context) error {
	username := c.Param("username")
	rank, _ := h.LeaderboardRepo.GetUserRank(context.TODO(), username)
	score, _ := h.LeaderboardRepo.GetUserScore(context.TODO(), username)
	fmt.Println(&leaderboard.User{Name: username, Score: score, Rank: rank})
	return c.JSON(http.StatusOK, &leaderboard.User{Name: username, Score: score, Rank: rank})
}
