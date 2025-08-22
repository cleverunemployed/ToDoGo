package main

import "github/cleverunemployed/ToDoGo/test"

// import (
// 	handlers "github/cleverunemployed/ToDoGo/internal/api"

// 	"github.com/gin-gonic/gin"
// )

func main() {
	// r := gin.Default()

	// api := r.Group("/api")
	// {
	// 	v1 := api.Group("/v1")
	// 	{
	// 		v1.GET("/get_all_tasks", handlers.GetAllTasks)
	// 		v1.GET("/get_task/:id", handlers.GetTask)
	// 		v1.POST("/add_task", handlers.AddTask)
	// 		v1.DELETE("/delete_task", handlers.DeleteTask)
	// 		v1.PATCH("/update_task", handlers.UpdateTask)
	// 	}
	// }
	// r.Run()
	test.TestMain()
}
