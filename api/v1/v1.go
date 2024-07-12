package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/zlalabs/bunny/api/v1/user"
)

func RouterV1(e *echo.Echo) {

	v1 := e.Group("/v1")
	user.Router(v1)
}
