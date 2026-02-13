package main

import (
	"context"
	"go_book_api/config"
	"go_book_api/internal/handler"
	"go_book_api/internal/infrastructure/repository"
	"go_book_api/internal/middleware"
	"go_book_api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	ctx := context.Background()

	userRepo := repository.NewUserRepository(ctx, config.DB)
	bookRepo := repository.NewBookRepository(ctx, config.DB)
	userService := services.NewUserService(ctx, userRepo, bookRepo)
	bookService := services.NewBookService(ctx, bookRepo)

	userHandler := handler.NewUserHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)

	//routes
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	users := protected.Group("/users")
	{
		users.GET("/:id", userHandler.GetUser)
		users.GET("/:id/books", userHandler.GetUserBooks)
	}

	books := protected.Group("/books")
	{
		books.GET("", bookHandler.GetBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	protected.POST("/logout", handler.Logout)

	r.Run(":8080")
}
