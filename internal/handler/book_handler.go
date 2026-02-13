package handler

import (
	"net/http"
	"strconv"

	"go_book_api/internal"
	"go_book_api/internal/services"
	"go_book_api/utils"

	"github.com/gin-gonic/gin"
)

type BookHanlder interface {
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *bookHandler {
	return &bookHandler{
		bookService: bookService,
	}
}

// list all books
func (h *bookHandler) GetBooks(c *gin.Context) {
	books := h.bookService.GetAll(c)
	utils.ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)
}

// get book by id
func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)
	book, err := h.bookService.GetById(c, uint(uintId))
	if err != nil {
		utils.ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest internal.BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	book, err := h.bookService.Create(c, bookRequest)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}

	utils.ResponseJSON(c, http.StatusCreated, "Book created successfully", book)
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)
	var bookRequest internal.BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		utils.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	_, err := h.bookService.GetById(c, uint(uintId))
	if err != nil {
		utils.ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	book, err := h.bookService.Update(c, bookRequest)
	if err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, "Failed to update book", nil)
		return
	}

	utils.ResponseJSON(c, http.StatusOK, "Book updated successfully", book)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.Atoi(id)
	if err := h.bookService.Delete(c, uint(uintId)); err != nil {
		utils.ResponseJSON(c, http.StatusInternalServerError, "Failed to delete book", nil)
		return
	}
	utils.ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}
