package main

// import "github/cleverunemployed/ToDoGo/test"

import (
	handlers "github/cleverunemployed/ToDoGo/internal/api"
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/middleware"
	"log"

	_ "github/cleverunemployed/ToDoGo/docs" // ОЧЕНЬ ВАЖНО: импортируйте сгенерированный docs package

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Awesome API
// @version 1.0
// @description This is a sample server for a awesome API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config, err := configs.Init()
	if err != nil {
		log.Fatal("Config isn't download!")
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
