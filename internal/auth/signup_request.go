package auth

import "go_book_api/internal/user"

type SignupRequest struct {
	Name     string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,min=6"`
}

func (r *SignupRequest) ToModel() user.User {
	return user.User{
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
	}
}
