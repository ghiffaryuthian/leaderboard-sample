package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MakeMiscHandler registers endpoints for healthz
func MakeMiscHandler(e *echo.Echo) {
	e.GET("/healthz", healthz)
}

func healthz(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
