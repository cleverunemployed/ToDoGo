package db

import (
	"context"
	"github/cleverunemployed/ToDoGo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (db *DBConnection) GetAllTasksForUser(userID string) ([]models.Tasks, error) {
	var tasks []models.Tasks
	var ctx = context.Background()
	err := db.DB.WithContext(ctx).Table("tasks").
		Joins("INNER JOIN user_task ON tasks.id = user_task.idtask").
		Where("user_task.iduser = ?", userID).
		Find(&tasks).Error
	return tasks, err
}

func (db *DBConnection) UpdateTask(taskData models.Tasks) error {
	var ctx = context.Background()
	err := db.DB.WithContext(ctx).Save(&taskData).Error
	return err
}

func (db *DBConnection) DeleteTask(taskID string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ", taskID).Delete(&models.Tasks{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id_task = ", taskID).Delete(&models.UserTask{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (db *DBConnection) CreateTask(taskData models.Tasks) (taskID uuid.UUID, err error) {
	var task = models.Tasks{
		Title:       taskData.Title,
		Description: taskData.Description,
		Completed:   taskData.Completed,
		DateStart:   taskData.DateStart,
		DateEnd:     taskData.DateEnd,
	}

	result := db.DB.Create(&task)

	return task.ID, result.Error
}
