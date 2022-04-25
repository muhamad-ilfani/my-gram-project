package config

import (
	"fmt"
	"my-gram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "April2022"
	dbname   = "mygram"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Media{})
	return db
}
