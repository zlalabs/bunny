package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbn := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pwd := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable TimeZone=Asia/Bangkok", host, port, dbn, user, pwd)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Connected db")

	DB = db
	migrate(db)
	return db, nil
}

func migrate(db *gorm.DB) {

	// envApp := os.Getenv("API_PORT")
	// if envApp == util.ENV_DEV {
	db.AutoMigrate(&User{})
	// }
}
