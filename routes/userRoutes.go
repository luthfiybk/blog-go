package routes

import (
	"blog-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	base := r.Group("/api")

	uc := base.Group("/users")

	uc.GET("", controllers.GetUsers)
	uc.POST("/signup", controllers.SignUp)
	uc.POST("/login", controllers.Login)
	uc.GET("/auth", controllers.Validate)
	uc.GET("/logout", controllers.Logout)
}