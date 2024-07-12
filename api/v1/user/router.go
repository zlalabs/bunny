package user

import (
	"github.com/labstack/echo/v4"
	"github.com/zlalabs/bunny/api/v1/user/handlers"
	"github.com/zlalabs/bunny/api/v1/user/repositories"
	"github.com/zlalabs/bunny/api/v1/user/services"
)

func Router(e *echo.Group) {

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	user := e.Group("/users")
	user.GET("", userHandler.GetAll)
	user.POST("", userHandler.Create)
}
