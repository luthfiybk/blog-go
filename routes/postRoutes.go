package routes

import (
	"blog-go/controllers"
	"blog-go/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {
	base := r.Group("/api")

	pc := base.Group("/posts")
	pc.GET("", controllers.GetPosts)
	pc.POST("", middlewares.RequireAuth, controllers.CreatePost)
	pc.GET("/:id", controllers.GetPost)
	pc.PUT("/:id", controllers.UpdatePost)
	pc.DELETE("/:id", controllers.DeletePost)
}