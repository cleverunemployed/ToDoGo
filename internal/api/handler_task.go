package api

import (
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/db"
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/schemas"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateTask создает новую задачу
// @Summary Создание задачи
// @Description Создает новую задачу для пользователя
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body schemas.CreateTaskRequest true "Данные задачи"
// @Success 201 {object} map[any]any
// @Failure 400 {object} map[any]any
// @Failure 500 {object} map[any]any
// @Router /api/v1/create/task [post]
func CreateTask(c *gin.Context) {
	var jsonData = schemas.CreateTaskRequest{}
	var taskData models.Tasks

	if err := c.BindJSON(&jsonData); err != nil {
		log.Println(err)
	} else {
		taskData.Title = jsonData.Title
		taskData.Description = jsonData.Description
		taskData.Completed = jsonData.Completed
		taskData.DateStart = jsonData.DateStart
		taskData.DateEnd = jsonData.DateEnd
	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	id, err := db.CreateTask(taskData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(201, gin.H{
		"msg":     "Task created",
		"id_user": id,
	})
}

// ReadAllTasks получает все задачи пользователя
// @Summary Получение всех задач
// @Description Возвращает список всех задач для указанного пользователя
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body schemas.ReadTaskRequest true "ID пользователя"
// @Success 200 {object} map[any]any
// @Failure 400 {object} map[any]any
// @Failure 500 {object} map[any]any
// @Router /api/v1/get/all/tasks [post]
func ReadAllTasks(c *gin.Context) {
	var jsonData = schemas.ReadTaskRequest{}
	c.BindJSON(&jsonData)

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	tasks, err := db.GetAllTasksForUser(jsonData.IDUser.String())
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg":   "get tasks",
		"tasks": tasks,
	})
}

// UpdateTask обновляет существующую задачу
// @Summary Обновление задачи
// @Description Обновляет данные существующей задачи
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body schemas.UpdateTaskRequest true "Обновленные данные задачи"
// @Success 200 {object} map[any]any
// @Failure 400 {object} map[any]any
// @Failure 500 {object} map[any]any
// @Router /api/v1/update/task [patch]
func UpdateTask(c *gin.Context) {
	var jsonData = schemas.UpdateTaskRequest{}
	var taskData models.Tasks

	if err := c.BindJSON(&jsonData); err != nil {
		log.Println(err)
	} else {
		taskData.Title = jsonData.Title
		taskData.Description = jsonData.Description
		taskData.Completed = jsonData.Completed
		taskData.DateStart = jsonData.DateStart
		taskData.DateEnd = jsonData.DateEnd
	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	err = db.UpdateTask(taskData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(201, gin.H{
		"msg": "Task Updated",
	})
}

// DeleteTask удаляет задачу
// @Summary Удаление задачи
// @Description Удаляет задачу по ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body schemas.DeleteTaskRequest true "ID задачи для удаления"
// @Success 200 {object} map[any]any
// @Failure 400 {object} map[any]any
// @Failure 500 {object} map[any]any
// @Router /api/v1/delete/task [delete]
func DeleteTask(c *gin.Context) {
	var jsonData = schemas.DeleteTaskRequest{}

	if err := c.BindJSON(&jsonData); err != nil {
		log.Println(err)
	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	err = db.DeleteTask(jsonData.ID.String())
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg": "Task Deleted",
	})
}
