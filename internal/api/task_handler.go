package api

import (
	"github/cleverunemployed/ToDoGo/internal/db"
	"github/cleverunemployed/ToDoGo/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var connection = db.Init("internal\\db\\db.json")

func AddTask(c *gin.Context) {
	var title string = c.Query("title")

	err := connection.AddTask(models.SchemaTask{
		Title: title,
	})

	if err != nil {
		c.JSON(500, gin.H{"message": "Server error", "error": err})
	}
	c.JSON(201, gin.H{"message": "success"})
}

func GetTask(c *gin.Context) {
	var id string = c.Param("id")
	var data models.Task

	data, err := connection.GetTask(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "Server error", "error": err})
	}
	c.JSON(200, gin.H{"message": "success", "data": data})
}

func GetAllTasks(c *gin.Context) {
	var data []models.Task

	data, err := connection.GetAllTasks()
	if err != nil {
		c.JSON(500, gin.H{"message": "Server error", "error": err})
	}
	c.JSON(200, gin.H{"message": "success", "data": data})
}

func UpdateTask(c *gin.Context) {
	var id string = c.Query("id")
	var completed bool
	completed, err := strconv.ParseBool(c.Query("completed"))

	if err != nil {
		c.JSON(400, gin.H{"message": "Params error", "error": err})
	}

	err = connection.UpdateTask(id, completed)
	if err != nil {
		c.JSON(500, gin.H{"message": "Server error", "error": err})
	}
	c.JSON(200, gin.H{"message": "success"})
}

func DeleteTask(c *gin.Context) {
	var id string = c.Query("id")

	err := connection.DeleteTask(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "Server error", "error": err})
	}
	c.JSON(200, gin.H{"message": "success"})
}
