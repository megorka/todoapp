package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/megorka/todoapp/authorization/models"
	"github.com/megorka/todoapp/authorization/pkg/jwt"
)

type Client struct {
    baseURL string
    client  *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 5 * time.Second,
        },
    }
}

func (c *Client) RegisterUser(ctx context.Context, username, email, password string) error {

	hashedPassword, err := jwt.HashPassword(password)
	if err != nil {
		return fmt.Errorf("HashPassword: %w", err)
	}

  body := map[string]string{
    "username": username,
    "email":    email,
    "password": hashedPassword,
  }
  jsonBody, err := json.Marshal(body)
  if err != nil {
    return fmt.Errorf("failed to marshal request body: %w", err)
  }
  req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8181/api/v1/auth/signup", bytes.NewBuffer(jsonBody))
  if err != nil {
    return fmt.Errorf("failed to create request: %w", err)
  }
  req.Header.Set("Content-Type", "application/json")

  resp, err := c.client.Do(req)
  if err != nil {
    return fmt.Errorf("failed to send request: %w", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    responseBody, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("user service returned status %d: %s", resp.StatusCode, string(responseBody))
  }

  return nil
}

func (c *Client) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	url := "http://localhost:8181/api/v1/user/" + email
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed create request: %w", err)
	}
	
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user service returned status %d: %s", resp.StatusCode, string(responseBody))
	}
	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("Failed to decode: %w", err)
	}

	return &user, nil
}