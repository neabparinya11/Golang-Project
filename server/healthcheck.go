package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	HealthCheck struct {
		App    string `json:"app"`
		Status string `json:"status"`
	}
)

// Check service is running
func (s *Server) HealthCheckService(c echo.Context) error {
	return response.SuccessResponse(c, http.StatusOK, &HealthCheck{
		App: s.cfg.App.Name,
		Status: "OK",
	})
}
