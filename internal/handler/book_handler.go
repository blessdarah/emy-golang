package api

import (
	"net/http"

	"go_book_api/config"
	"go_book_api/internal/model"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book model.Book

	// bind the request body to the book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	config.DB.Create(&book)
	model.ResponseJSON(c, http.StatusCreated, "Book created successfully", book)
}

func GetBooks(c *gin.Context) {
	var books []model.Book

	config.DB.Find(&books)
	model.ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)

}

func GetBook(c *gin.Context) {
	var book model.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}

func UpdateBook(c *gin.Context) {
	var book model.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// bind the request body
	if err := c.ShouldBindJSON(&book); err != nil {
		model.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	config.DB.Save(&book)
	model.ResponseJSON(c, http.StatusOK, "Book updated successfully", book)
}

func DeleteBook(c *gin.Context) {
	var book model.Book
	if err := config.DB.Delete(&book, c.Param("id")).Error; err != nil {
		model.ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	model.ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}
