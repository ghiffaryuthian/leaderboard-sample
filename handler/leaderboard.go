package handler

import (
	"context"
	"net/http"
	"strconv"

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
	e.POST("/leaderboards/ranks", h.rankUser)
}

func (h *lbHandler) getUserDetails(c echo.Context) error {
	username := c.Param("username")
	user, err := h.LeaderboardService.GetMemberDetails(context.TODO(), username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *lbHandler) rankUser(c echo.Context) error {
	p := echo.Map{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	username := p["username"].(string)
	score, err := strconv.ParseFloat(p["score"].(string), 64)
	if err != nil {
		return err
	}
	user, err := h.LeaderboardService.RankMember(context.TODO(), username, score)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
