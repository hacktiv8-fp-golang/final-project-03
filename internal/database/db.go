package database

import (
	"final-project-03/internal/config"
	"final-project-03/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := config.GetDBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(model.User{}, model.Category{}, model.Task{})
}

func GetDB() *gorm.DB {
	return db
}
