package auth

import (
	"context"
	"go_book_api/internal/user"
	"log/slog"
)

type userStore interface {
	GetByEmail(c context.Context, email string) (*user.User, error)
	Create(c context.Context, user user.User) (*user.User, error)
}

type authService struct {
	ctx    context.Context
	users  userStore
	logger *slog.Logger
}

func NewAuthService(ctx context.Context, store userStore, l *slog.Logger) *authService {
	return &authService{
		ctx:    ctx,
		users:  store,
		logger: l,
	}
}

// Login authenticates a user
func (s *authService) LoginWithEmail(email string) *user.User {
	u, _ := s.users.GetByEmail(s.ctx, email)
	return u
}

func (s *authService) Signup(req SignupRequest) (*user.User, error) {
	return s.users.Create(s.ctx, req.ToModel())
}
