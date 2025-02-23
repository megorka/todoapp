package service

import (
	"context"
	"fmt"
	"github.com/megorka/todoapp/authorization/adapters"

	"github.com/megorka/todoapp/authorization/pkg/jwt"
)

type Service struct {
	adapters *adapters.Adapter
}

func NewService(adapters *adapters.Adapter) *Service {
	return &Service{adapters: adapters}
}

func (s *Service) RegisterUser(ctx context.Context, username, email, password string) error {
	err := s.adapters.RegisterUser(ctx, username, email, password)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	fmt.Printf("User registered successfully: %s\n", email)
	return nil
}

func (s *Service) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.adapters.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !jwt.CheckPasswordHash(password, user.Password) {
		return "", err
	}

	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
