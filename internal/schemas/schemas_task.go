package schemas

import (
	"time"

	"github.com/google/uuid"
)

// CreateTaskRequest DTO для создания задачи
// @Description Запрос на создание новой задачи
type CreateTaskRequest struct {
	Title       string    `json:"title" example:"Завершить проект"`                              // Название задачи
	Description string    `json:"description" example:"Необходимо завершить все задачи проекта"` // Описание задачи
	Completed   bool      `json:"completed" example:"false"`                                     // Статус выполнения задачи
	DateStart   time.Time `json:"date_start" example:"2024-01-15T10:00:00Z"`                     // Дата начала выполнения
	DateEnd     time.Time `json:"date_end" example:"2024-01-20T18:00:00Z"`                       // Дата окончания выполнения
}

// TaskResponse DTO для ответа с данными задачи
// @Description Ответ с информацией о задаче
type TaskResponse struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`             // UUID идентификатор задачи
	Title       string    `json:"title" example:"Завершить проект"`                              // Название задачи
	Description string    `json:"description" example:"Необходимо завершить все задачи проекта"` // Описание задачи
	Completed   bool      `json:"completed" example:"false"`                                     // Статус выполнения задачи
	DateStart   time.Time `json:"date_start" example:"2024-01-15T10:00:00Z"`                     // Дата начала выполнения
	DateEnd     time.Time `json:"date_end" example:"2024-01-20T18:00:00Z"`                       // Дата окончания выполнения
	CreatedAt   time.Time `json:"created_at" example:"2024-01-15T09:00:00Z"`                     // Дата создания записи
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-15T09:30:00Z"`                     // Дата последнего обновления
}

// ReadTaskRequest DTO для чтения задачи
// @Description Запрос на получение задачи по идентификатору пользователя
type ReadTaskRequest struct {
	IDUser uuid.UUID `json:"id_user" example:"550e8400-e29b-41d4-a716-446655440000"` // UUID идентификатор пользователя
}

// UpdateTaskRequest DTO для обновления задачи
// @Description Запрос на обновление существующей задачи
type UpdateTaskRequest struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"` // UUID идентификатор задачи
	Title       string    `json:"title" example:"Обновленное название задачи"`       // Новое название задачи
	Description string    `json:"description" example:"Обновленное описание задачи"` // Новое описание задачи
	Completed   bool      `json:"completed" example:"true"`                          // Новый статус выполнения
	DateStart   time.Time `json:"date_start" example:"2024-01-16T10:00:00Z"`         // Новая дата начала
	DateEnd     time.Time `json:"date_end" example:"2024-01-25T18:00:00Z"`           // Новая дата окончания
}

// DeleteTaskRequest DTO для удаления задачи
// @Description Запрос на удаление задачи по идентификатору
type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"` // UUID идентификатор задачи для удаления
}
