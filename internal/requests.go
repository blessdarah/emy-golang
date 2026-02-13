package internal

import model "go_book_api/internal/infrastructure/gorm"

type UserRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (r UserRequest) ToModel() model.User {
	return model.User{
		Username: r.Username,
		Password: r.Password,
	}
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (r BookRequest) ToModel() model.Book {
	return model.Book{
		Title:  r.Title,
		Author: r.Author,
		Year:   r.Year,
	}
}
