package repositories

import (
	"github.com/zlalabs/bunny/internal/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]database.User, error)
	Create(data database.User) (*database.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &repository{
		db: database.DB,
	}
}

func (r *repository) Create(data database.User) (*database.User, error) {

	q := r.db.Create(&data)
	if q.Error != nil {
		return nil, q.Error
	}
	return &data, nil
}

func (r *repository) GetAll() ([]database.User, error) {
	users := make([]database.User, 0)

	q := r.db.Find(&users)
	if q.Error != nil {
		return nil, q.Error
	}
	return users, nil
}
