package users

import (
	"context"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, req CreateUserRequest,
) (*UserResponse, error) {
	user := User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
	}

	err := s.repo.Create(ctx, &user)

	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
