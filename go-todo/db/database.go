package db

import (
	"fmt"
	"go-todo/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		fmt.Println("Migration ran successfully!")
	}

	fmt.Println("Database created and migrated successfully")
	return DB, err
}
