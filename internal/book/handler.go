package book

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService BookService
	logger      *slog.Logger
}

type BookService interface {
	GetAll(c *gin.Context) []Book
	GetById(c *gin.Context, id uint) (*Book, error)
	Create(c *gin.Context, br CreateRequest) (*Book, error)
	Update(c *gin.Context, br UpdateRequest) (*Book, error)
	Delete(c *gin.Context, id uint) error
}

func NewBookHandler(bs BookService, l *slog.Logger) *bookHandler {
	return &bookHandler{
		bookService: bs,
		logger:      l,
	}
}

func (h *bookHandler) Index(c *gin.Context) {
	books := h.bookService.GetAll(c)

	// convert books to collection
	collection := make([]Response, len(books))
	for i, b := range books {
		collection[i] = b.ToResponse()
	}
	c.JSON(http.StatusOK, gin.H{
		"data": collection,
	})
}

// Show returns a single book by id
func (h *bookHandler) Show(c *gin.Context) {
	id := c.Param("id")

	// convert string id to uint
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		h.logger.Error("convert id to uint",
			slog.String("action", "show book"),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id %s in path", id),
		})
		return
	}

	book, err := h.bookService.GetById(c, uint(uintId))
	if err != nil {
		h.logger.Error("get book",
			slog.String("action", "show book"),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book.ToResponse(),
	})

}

// CreateBook creates a new book
// sample request:
//
//	{
//	  "title": "The Lord of the Rings",
//	  "author": "J.R.R. Tolkien",
//	  "year": 1954
//	}
func (h *bookHandler) Create(c *gin.Context) {
	var bookRequest CreateRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		h.logger.Error(
			"book validation error",
			slog.String("action", "create book"),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	book, err := h.bookService.Create(c, bookRequest)
	if err != nil {
		h.logger.Error(
			"create book",
			slog.String("error", err.Error()),
			slog.Any("payload", bookRequest),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create book",
		})
		return
	}

	h.logger.Info("book created",
		slog.Uint64("id", uint64(book.ID)),
	)
	c.JSON(http.StatusCreated, gin.H{
		"data": book.ToResponse(),
	})
}

// UpdateBook updates a book
// sample request:
//
//	{
//	  "title": "The Lord of the Rings",
//	  "author": "J.R.R. Tolkien",
//	  "year": 1954
//	}
func (h *bookHandler) Update(c *gin.Context) {
	var bookRequest UpdateRequest

	// get book id from path
	strId := c.Param("id")

	// convert string id to uint
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		h.logger.Error("convert id to uint",
			slog.String("action", "update book"),
			slog.Uint64("id", id),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id %s in path", strId),
		})
		return
	}

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		h.logger.Error("validate book",
			slog.String("action", "update book"),
			slog.Any("payload", bookRequest),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = h.bookService.GetById(c, uint(id))
	if err != nil {
		h.logger.Error("get book",
			slog.String("action", "update book"),
			slog.Uint64("id", id),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}

	book, err := h.bookService.Update(c, bookRequest)
	if err != nil {
		h.logger.Error(
			"update book",
			slog.String("error", err.Error()),
			slog.Any("payload", bookRequest),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update book",
		})
		return
	}

	h.logger.Info("book updated",
		slog.Uint64("id", uint64(book.ID)),
	)
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// Delete deletes a book by id
func (h *bookHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Error("convert id to uint",
			slog.String("action", "delete book"),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid id %s in path", id),
		})
		return
	}

	if err := h.bookService.Delete(c, uint(uintId)); err != nil {
		h.logger.Error("delete book",
			slog.String("action", "delete book"),
			slog.String("error", err.Error()),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete book",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
