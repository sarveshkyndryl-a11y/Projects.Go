package models

import "time"

type AdminSecurity struct {
	UserID            int64     `json:"user_id"`
	TwoFactorEnabled  bool      `json:"two_factor_enabled"`
	TwoFactorType     string    `json:"two_factor_type"` // TOTP, EMAIL
	TwoFactorSecret   string    `json:"-"`               // encrypted
	AllowedIPRanges   string    `json:"allowed_ip_ranges,omitempty"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}