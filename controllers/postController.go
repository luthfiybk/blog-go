package controllers

import (
	"blog-go/inits"
	"blog-go/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title	string
		Body	string
		Likes	int
		Draft	bool
		Author	string
		UserID 	uint 	`json:"user_id"`
	}

	ctx.BindJSON(&body)

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message": "Failed to get user!",
			"error": "User not found!",
		})
	}
	body.UserID = user.(models.User).ID

	post := models.Post{
		Title: body.Title,
		Body: body.Body,
		Likes: body.Likes,
		Draft: body.Draft,
		Author: body.Author,
		UserID: body.UserID,
	}

	fmt.Println(post)

	result := inits.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message": "Failed to create post!",
			"error": result.Error,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"statusCode": 201,
		"message": "Post created successfully!",
		"data": post,
	})
}

func GetPosts(ctx *gin.Context) {
	var posts []models.Post

	result := inits.DB.Find(&posts)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message": "Failed to fetch posts!",
			"error": result.Error,
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Posts fetched successfully!",
		"data": posts,
	})
}

func GetPost(ctx *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))

	if result.Error != nil{
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message": "Post not found!",
			"error": result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Post fetched successfully!",
		"data": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))

	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message": "Post not found!",
			"error": result.Error,
		})
		return
	}

	var body struct {
		Title	string
		Body	string
		Likes	int
		Draft	bool
		Author	string
	}

	ctx.BindJSON(&body)

	post.Title = body.Title
	post.Body = body.Body
	post.Likes = body.Likes
	post.Draft = body.Draft
	post.Author = body.Author

	result = inits.DB.Save(&post)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message": "Failed to update post!",
			"error": result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Post updated successfully!",
		"data": post,
	})
}

func DeletePost(ctx *gin.Context) {
	var post models.Post

	found := inits.DB.First(&post, ctx.Param("id"))

	if found.Error != nil {
		ctx.JSON(400, gin.H{
			"statusCode": 400,
			"message": "Post not found!",
			"error": found.Error,
		})
		return
	}

	result := inits.DB.Delete(&post)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"statusCode": 500,
			"message": "Failed to delete post!",
			"error": result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message": "Post deleted successfully!",
		"data": post,
	})
}