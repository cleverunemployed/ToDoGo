package db

import (
	"github/cleverunemployed/ToDoGo/internal/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	DB *gorm.DB
}

func ConnectDB(db_url string) (*DBConnection, error) {
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
		return nil, err
	}

	return &DBConnection{DB: db}, nil
}

func (db DBConnection) MigrateDB() error {
	return db.DB.AutoMigrate(&models.Tasks{}, &models.Users{}, &models.UserTask{})
}
