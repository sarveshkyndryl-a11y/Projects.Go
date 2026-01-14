package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Contact_email string    `json:"contact_email"`
	Contact_phone string    `json:"contact_phone"`
	CreatedAt    time.Time `json:"created_at"`
}

