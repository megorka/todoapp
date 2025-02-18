package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/megorka/todoapp/user_service/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUsers() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := r.db.Exec(query)
	if err != nil {
		log.Printf("Failed to create users table: %v", err)
		return err
	}
	log.Println("Successfully created users table or already exists")
	return nil
}

func (repo *Repository) CreateUser(username, email, password string) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := repo.db.Exec(query, username, email, password)
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	if err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username,&user.Email, &user.Password,); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}