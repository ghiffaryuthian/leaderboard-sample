package handler

import (
	"context"
	"net/http"

	"github.com/ghiffaryuthian/leaderboard-sample/leaderboard"
	"github.com/labstack/echo/v4"
)

type lbHandler struct {
	LeaderboardService leaderboard.Service
}

// MakeLeaderboardHandler registers endpoints for healthz
func MakeLeaderboardHandler(e *echo.Echo, lbService leaderboard.Service) {
	h := lbHandler{LeaderboardService: lbService}
	e.GET("/leaderboards/ranks/:username", h.getUserDetails)
}

func (h *lbHandler) getUserDetails(c echo.Context) error {
	return c.JSON(http.StatusOK, h.LeaderboardService.GetMemberDetails(context.TODO(), c.Param("username")))
}
