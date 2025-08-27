package api

import (
	"github/cleverunemployed/ToDoGo/internal/configs"
	"github/cleverunemployed/ToDoGo/internal/db"
	"github/cleverunemployed/ToDoGo/internal/models"
	"github/cleverunemployed/ToDoGo/internal/schemas"
	"github/cleverunemployed/ToDoGo/internal/tools"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	token, err := tools.NewAccessToken(tools.UserClaims{
		ID:    id.String(),
		Email: jsonData.Email,
	})
	if err != nil {
		log.Fatalln("Jwt isn't created!")
	}

	refreshToken, err := tools.NewRefreshToken(jwt.RegisteredClaims{
		Subject:   id.String(), // Usually the user ID
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)), // 7 days for refresh token
		Issuer:    "todoapp",                                              // Your application name
		// Add other standard JWT claims as needed
		// Audience:  []string{"your-audience"},
		// ID:        "unique-token-id",
	})
	if err != nil {
		log.Fatalln("Refresh Jwt isn't created!")
	}

	c.JSON(201, gin.H{
		"msg": "User created",
		"data": map[any]any{
			"token":         token,
			"refresh_token": refreshToken,
			"id_user":       id,
		},
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

	user, err := db.ReadUser(jsonData)
	if err != nil {
		log.Println(err)
	}
	user.Password = "[secret]"

	token, err := tools.NewAccessToken(tools.UserClaims{
		ID:    user.ID.String(),
		Email: jsonData.Email,
	})
	if err != nil {
		log.Fatalln("Jwt isn't created!")
	}

	refreshToken, err := tools.NewRefreshToken(jwt.RegisteredClaims{
		Subject:   user.ID.String(), // Usually the user ID
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)), // 7 days for refresh token
		Issuer:    "todoapp",                                              // Your application name
		// Add other standard JWT claims as needed
		// Audience:  []string{"your-audience"},
		// ID:        "unique-token-id",
	})
	if err != nil {
		log.Fatalln("Refresh Jwt isn't created!")
	}

	c.JSON(200, gin.H{
		"msg": "User is entered",
		"data": map[any]any{
			"token":         token,
			"refresh_token": refreshToken,
			"user_data":     user,
		},
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
