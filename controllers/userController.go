package controllers

import (
	"blog-go/inits"
	"blog-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message":    "Failed to generate token!",
			"error":      err,
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "localhost", false, true)
}

func GetUsers(ctx *gin.Context) {
	var users []models.User

	err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message":    "Failed to fetch users!",
			"error":      err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Users fetched successfully!",
		"data":       users,
	})
}

func Validate(ctx *gin.Context) {
	user, err := ctx.Get("user")

	if err != false {
		ctx.JSON(401, gin.H{
			"statusCode": 401,
			"message":    "User is not authenticated!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "User is authenticated!",
		"data":       user,
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "User logged out successfully!",
	})
}