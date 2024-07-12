package api

import (
	"github.com/labstack/echo/v4"
	v1 "github.com/zlalabs/bunny/api/v1"
)

func Server(e *echo.Echo) {
	v1.RouterV1(e)
}
