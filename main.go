package main

import (
	"blog-go/inits"

	"github.com/gin-gonic/gin"

	"blog-go/routes"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	routes.PostRouter(r)

	r.Run(":5000")
}