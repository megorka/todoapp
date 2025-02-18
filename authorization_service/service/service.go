package service

import (
	"context"
	"fmt"

	"github.com/megorka/todoapp/authorization/client"
	"github.com/megorka/todoapp/authorization/pkg/jwt"
)

type Service struct {
  client *client.Client
}

func NewService(client *client.Client) *Service {
  return &Service{client: client}
}

func (s *Service) RegisterUser(ctx context.Context, username, email, password string) error {
  err := s.client.RegisterUser(ctx, username, email, password)
  if err != nil {
    return fmt.Errorf("failed to register user: %w", err)
  }
  fmt.Printf("User registered successfully: %s\n", email)
  return nil
}


func (s *Service) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.client.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("Tyt")
		fmt.Println()
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