package models

import (
	"time"
)

type SupportEngineerProfile struct {
	UserID        int64     `json:"user_id"`
	EmployeeID    string    `json:"employee_id"`
	Department    string    `json:"department"`
	ExpertiseArea string    `json:"expertise_area"`
	IsAvailable   bool      `json:"is_available"`
	CreatedAt     time.Time `json:"created_at"`
}