package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zlalabs/bunny/api/v1/user/commands"
	"github.com/zlalabs/bunny/api/v1/user/services"
)

type IUserHandler interface {
	GetAll(c echo.Context) error
	Create(c echo.Context) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) IUserHandler {
	return &userHandler{
		userService: userService,
	}
}

// Handler
func (h *userHandler) GetAll(c echo.Context) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusGone, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) Create(c echo.Context) error {
	u := new(commands.CreateUserRequest)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, _ := h.userService.Create(*u)

	return c.JSON(http.StatusOK, user)
}
