package routes

import (
	"blog-go/controllers"

	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {
	base := r.Group("/api")

	pc := base.Group("/posts")
	pc.GET("", controllers.GetPosts)
	pc.POST("", controllers.CreatePost)
	pc.GET("/:id", controllers.GetPost)
	pc.PUT("/:id", controllers.UpdatePost)
	pc.DELETE("/:id", controllers.DeletePost)
}