package main

import (
	"go_book_api/config"
	"go_book_api/internal/auth"
	api "go_book_api/internal/handler"
	"go_book_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	//routes
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/book", api.CreateBook)
	protected.GET("/books", api.GetBooks)
	protected.GET("/book/:id", api.GetBook)
	protected.PUT("/book/:id", api.UpdateBook)
	protected.DELETE("/book/:id", api.DeleteBook)
	protected.POST("/logout", auth.Logout)

	r.Run(":8080")
}
