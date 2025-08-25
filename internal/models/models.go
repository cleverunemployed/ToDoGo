package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string
	Description string
	Completed   bool
	DateStart   time.Time
	DateEnd     time.Time
}

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email    string    `gorm:"unique"`
	Password string
}

type UserTask struct {
	gorm.Model
	IDUser uuid.UUID
	IDTask uuid.UUID
}
