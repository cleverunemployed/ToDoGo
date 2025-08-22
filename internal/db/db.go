package db

import (
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/tools"
)

type DBConnection struct {
	url string
}

func Init(url string) *DBConnection {
	return &DBConnection{
		url: url,
	}
}

func (db DBConnection) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	tasks, err := tools.LoadTasks(db.url)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (db DBConnection) GetTask(id string) (models.Task, error) {
	var task models.Task
	var tasks []models.Task

	tasks, err := tools.LoadTasks(db.url)
	if err != nil {
		return task, err
	}

	for _, value := range tasks {
		if id == value.ID {
			return task, nil
		}
	}

	return task, nil
}

func (db DBConnection) AddTask(data models.SchemaTask) error {
	var tasks []models.Task

	tasks, err := tools.LoadTasks(db.url)
	if err != nil {
		return err
	}

	tasks[len(tasks)] = models.Task{
		ID:       tools.GetRandomString(),
		Title:    data.Title,
		Comleted: false,
	}

	tools.SaveTasks(db.url, tasks)
	return nil
}

func (db DBConnection) DeleteTask(id string) error {
	var tasks []models.Task

	tasks, err := tools.LoadTasks(db.url)
	if err != nil {
		return err
	}
	for index, value := range tasks {
		if value.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
		}
	}

	return tools.SaveTasks(db.url, tasks)
}

func (db DBConnection) UpdateTask(id string, competed bool) error {
	var tasks []models.Task

	tasks, err := tools.LoadTasks(db.url)
	if err != nil {
		return err
	}
	for index, value := range tasks {
		if value.ID == id {
			tasks[index] = models.Task{
				ID:       tasks[index].ID,
				Title:    tasks[index].Title,
				Comleted: competed,
			}
		}
	}

	return tools.SaveTasks(db.url, tasks)
}
