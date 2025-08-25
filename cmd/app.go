package main

// import "github/cleverunemployed/ToDoGo/test"

import (
	handlers "github/cleverunemployed/ToDoGo/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/signin", handlers.ReadUser)
			v1.POST("/signup", handlers.CreateUser)
			v1.PATCH("/change/password", handlers.UpdatePassword)
			v1.PATCH("/change/email", handlers.UpdateEmail)
			v1.DELETE("/delete/user", handlers.DeleteUser)

			v1.POST("/create/task", handlers.CreateTask)
			v1.POST("/get/all/tasks", handlers.ReadAllTasks)
			v1.DELETE("/delete/task", handlers.DeleteTask)
			v1.PATCH("/update/task", handlers.UpdateTask)
		}
	}

	r.Run()
}
