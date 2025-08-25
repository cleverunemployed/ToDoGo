package api

import (
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/db"
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/schemas"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	var jsonData = schemas.CreateUserRequest{}
	c.BindJSON(&jsonData)

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	id, err := db.CreateUser(jsonData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(201, gin.H{
		"msg":     "User created",
		"id_user": id,
	})
}

func ReadUser(c *gin.Context) {
	var jsonData = schemas.CreateUserRequest{}
	c.BindJSON(&jsonData)

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	id, err := db.ReadUser(jsonData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg":     "User is entered",
		"id_user": id,
	})
}

func UpdatePassword(c *gin.Context) {
	var jsonData = schemas.UpdatePasswordRequest{}
	var userUpdateData models.Users

	err := c.BindJSON(&jsonData)
	if err != nil {
		log.Println(err)
	} else {
		userUpdateData.ID, err = uuid.Parse(jsonData.ID)

		if err != nil {
			log.Println(err)
		}
		userUpdateData.Password = jsonData.Password

	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	err = db.UpdateUser(userUpdateData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg": "User is updated",
	})
}

func UpdateEmail(c *gin.Context) {
	var jsonData = schemas.UpdateEmailRequest{}
	var userUpdateData models.Users

	err := c.BindJSON(&jsonData)
	if err != nil {
		log.Println(err)
	} else {
		userUpdateData.ID, err = uuid.Parse(jsonData.ID)

		if err != nil {
			log.Println(err)
		}
		userUpdateData.Email = jsonData.Email

	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	err = db.UpdateUser(userUpdateData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg": "User is updated",
	})
}

func DeleteUser(c *gin.Context) {
	var jsonData = schemas.DeleteUserRequest{}
	var userUpdateData models.Users

	err := c.BindJSON(&jsonData)
	if err != nil {
		log.Println(err)
	} else {
		userUpdateData.Email = jsonData.Email
		userUpdateData.Password = jsonData.Password
	}

	config, err := configs.Init()
	if err != nil {
		log.Println(err)
	}

	db, err := db.ConnectDB(config.Url)
	if err != nil {
		log.Println(err)
	}

	err = db.DeleteUser(userUpdateData)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"msg": "User is updated",
	})
}
