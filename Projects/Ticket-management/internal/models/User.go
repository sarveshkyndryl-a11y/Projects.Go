package models

import "time"

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`    // never expose
	Role         string `json:"role"` // ADMIN, SUPPORT, CUSTOMER

	IsActive          bool       `json:"is_active"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	FailedLoginCount  int        `json:"failed_login_count"`
	AccountLockedTill *time.Time `json:"account_locked_till,omitempty"`

	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}