package models

import "time"

type CustomerProfile struct {
	UserID      int64     `json:"user_id"`
	CompanyID   int64     `json:"company_id"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}