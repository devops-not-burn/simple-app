package redis

import (
	"github.com/labstack/echo/v4"
	"net/http"
)


func Getfromredis(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Puttoredis(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}