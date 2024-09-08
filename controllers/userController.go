package controllers

import (
	"blog-go/inits"
	"blog-go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var body struct {
		Name 		string
		Email    	string
		Password 	string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message":    "Invalid request body!",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message":    "Failed to hash password!",
			"error":      err,
		})
		return
	}

	user := models.User{
		Name: body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	result := inits.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message":    "Failed to create user!",
			"error":      result.Error,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"statusCode": 201,
		"message":    "User created successfully!",
		"data":       user,
	})
}

func Login(ctx *gin.Context){
	var body struct {
		Email		string
		Password 	string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message":    "Invalid request body!",
		})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message":    "User not found!",
			"error":      result.Error,
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message":    "Invalid password!",
		})
		return
	}

}