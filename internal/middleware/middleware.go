package middleware

import (
	"go_book_api/internal/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			model.ResponseJSON(c, http.StatusUnauthorized, "Missing token", nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			c.Next()
		} else {
			model.ResponseJSON(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
		}
	}
}

func Protected(c *gin.Context) {
	username, _ := c.Get("username")
	model.ResponseJSON(c, http.StatusOK, "Welcome "+username.(string), username)
}
