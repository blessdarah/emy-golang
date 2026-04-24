package auth

import "time"

type LoginResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
