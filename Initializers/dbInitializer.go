package Initializers

import (
	"ToDoList/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DBInitializer() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
}

func DBMigration() {
	DB.AutoMigrate(&models.Task{})
}
