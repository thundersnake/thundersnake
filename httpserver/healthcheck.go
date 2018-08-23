package httpserver

import (
	"github.com/labstack/echo"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func httpHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, healthCheckResponse{
		Status: "UP",
	})
}
