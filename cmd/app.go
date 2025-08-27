package main

// import "github/cleverunemployed/ToDoGo/test"

import (
	handlers "github/cleverunemployed/ToDoGo/internal/api"
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config, err := configs.Init()
	if err != nil {
		log.Fatal("Config isn't download!")
	}

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/signin", handlers.ReadUser)
			v1.POST("/signup", handlers.CreateUser)

			auth := v1.Group("/", middleware.AuthMiddleware(config.Secret))
			{
				auth.PATCH("/change/password", handlers.UpdatePassword)
				auth.PATCH("/change/email", handlers.UpdateEmail)
				auth.DELETE("/delete/user", handlers.DeleteUser)

				auth.POST("/create/task", handlers.CreateTask)
				auth.POST("/get/all/tasks", handlers.ReadAllTasks)
				auth.DELETE("/delete/task", handlers.DeleteTask)
				auth.PATCH("/update/task", handlers.UpdateTask)
			}
		}
	}

	r.Run()
}
