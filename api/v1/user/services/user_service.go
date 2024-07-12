package services

import (
	"github.com/zlalabs/bunny/api/v1/user/commands"
	"github.com/zlalabs/bunny/api/v1/user/repositories"
	"github.com/zlalabs/bunny/internal/database"
)

type UserService interface {
	GetAll() ([]database.User, error)
	Create(data commands.CreateUserRequest) (*database.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) GetAll() ([]database.User, error) {
	return service.userRepository.GetAll()
}

func (service *userService) Create(data commands.CreateUserRequest) (*database.User, error) {
	c := database.User{
		Username:  data.Username,
		Password:  data.Password,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}
	return service.userRepository.Create(c)
}
