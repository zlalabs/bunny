package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey" json:"id"`
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Activated bool
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()

	// if u.Role == "admin" {
	//   return errors.New("invalid role")
	// }
	return
}
