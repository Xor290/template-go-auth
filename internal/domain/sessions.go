package domain

import (
	"errors"
	"time"
)

type Session struct {
	SessionID    string    `json:"session_id"`
	UserID       uint      `json:"user_id"`
	Email        string    `json:"email"`
	TokenFamily  string    `json:"token_family"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

var (
	ErrSessionNotFound = NewError("session not found")
	ErrSessionExpired  = NewError("session expired")
)

func NewError(message string) error {
	return errors.New(message)
}
