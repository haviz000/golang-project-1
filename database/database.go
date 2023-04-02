package database

import (
	"fmt"
	"log"

	"github.com/haviz000/golang-project-1/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "jaka"
	dbName   = "book_project_1"
	port     = 5432
	db       *gorm.DB
	err      error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(entity.Book{})
}

func GetDB() *gorm.DB {
	return db
}
