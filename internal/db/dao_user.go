package db

import (
	"context"
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/schemas"
	"github/cleverunemployed/ToDoGo/internal/tools"
	"log"

	"github.com/google/uuid"
)

func (db *DBConnection) CreateUser(userData schemas.CreateUserRequest) (userID uuid.UUID, err error) {
	hashPassword, err := tools.HashPassword(userData.Password)
	if err != nil {
		log.Fatal("Hashing password don't work!")
	}
	var user = models.Users{
		Email:    userData.Email,
		Password: hashPassword,
	}

	result := db.DB.Create(&user)

	return user.ID, result.Error
}

func (db *DBConnection) UpdateUser(userData models.Users) error {
	var ctx = context.Background()
	err := db.DB.WithContext(ctx).Save(&userData).Error
	return err
}

func (db *DBConnection) ReadUser(userData schemas.CreateUserRequest) (models.Users, error) {
	var user models.Users
	var ctx = context.Background()

	hashPassword, err := tools.HashPassword(userData.Password)
	if err != nil {
		log.Fatal("Hashing password don't work!")
	}

	err = db.DB.WithContext(ctx).Table("users").
		Where("users.email = ?", userData.Email).
		Where("users.password = ?", hashPassword).
		Find(&user).Error
	return user, err
}

func (db *DBConnection) DeleteUser(userData models.Users) error {

	hashPassword, err := tools.HashPassword(userData.Password)
	if err != nil {
		log.Fatal("Hashing password don't work!")
		return err
	}

	return db.DB.
		Where("email = ", userData.Email).
		Where("password = ", hashPassword).
		Delete(&models.Users{}).Error
}
