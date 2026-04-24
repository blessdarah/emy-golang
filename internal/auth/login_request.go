package auth

type LoginRequest struct {
	Email    string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}
