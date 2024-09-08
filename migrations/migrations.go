package main

import (
	"blog-go/inits"
	"blog-go/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
	inits.DB.AutoMigrate(&models.User{})
}