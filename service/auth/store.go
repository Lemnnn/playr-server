package auth

import (
	"database/sql"
	"fmt"
	"playr-server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user *types.User) error {
	existsQuery := `
		SELECT 1 FROM users WHERE email = $1
	`
	var exists bool
	err := s.db.QueryRow(existsQuery, user.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil
	}

	insertQuery := `
		INSERT INTO users (email, first_name, last_name, avatar_url, created_at)
		VALUES ($1, $2, $3, $4, NOW())
	`
	_, err = s.db.Exec(insertQuery, user.Email, user.FirstName, user.LastName, user.AvatarURL)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
