package schemas

import (
	"time"

	"github.com/google/uuid"
)

// DTO для создания пользователя
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// DTO для ответа
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

// DTO для обновления пароля пользователя
type UpdatePasswordRequest struct {
	ID       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// DTO для обновления пароля пользователя
type UpdateEmailRequest struct {
	ID    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// DTO для обновления пароля пользователя
type DeleteUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
