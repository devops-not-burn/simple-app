package redis

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)


func Getfromredis(c echo.Context) error {
	if os.Getenv("REDIS_ENABLED") == "true" {
		return c.String(http.StatusOK, "OK")
	} else {
		return c.String(http.StatusOK, "Redis is not enabled")
	}

}

func Puttoredis(c echo.Context) error {
	if os.Getenv("REDIS_ENABLED") == "true" {
		return c.String(http.StatusOK, "OK")
	} else {
		return c.String(http.StatusOK, "Redis is not enabled")
	}
}