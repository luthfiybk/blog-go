package middlewares

import (
	"fmt"
	"blog-go/inits"
	"blog-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.JSON(401, gin.H{
			"statusCode": 401,
			"message":    "Unauthorized!",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "Token expired!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
	
		inits.DB.First(&user, int(claims["id"].(float64)))

		if user.ID == 0 {
			ctx.JSON(401, gin.H{
				"statusCode": 401,
				"message":    "User not found!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	
	ctx.Next()
}