package api

import (
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/db"
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/schemas"
	"log"

	"github.com/gin-gonic/gin"
)

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
