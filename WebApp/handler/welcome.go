package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Welcome (c echo.Context) error{
	return c.String(http.StatusOK,"welcome to My Go App ")
}