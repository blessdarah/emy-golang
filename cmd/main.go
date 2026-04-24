package main

import (
	"context"
	"go_book_api/config"
	"go_book_api/internal/book"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	var lev slog.Level
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     lev,
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, &opts))

	ctx := context.Background()

	bookRepo := book.NewBookRepository(ctx, config.DB)
	bookServ := book.NewBookService(ctx, bookRepo)

	bookHandler := book.NewBookHandler(bookServ, l)

	protected := r.Group("/api")
	// protected.Use(middleware.AuthMiddleware())

	books := protected.Group("/books")
	{
		books.GET("", bookHandler.Index)
		books.POST("", bookHandler.Create)
		books.GET("/:id", bookHandler.Show)
		books.PATCH("/:id", bookHandler.Update)
		books.DELETE("/:id", bookHandler.Delete)
	}

	r.Run(":8080")
}
