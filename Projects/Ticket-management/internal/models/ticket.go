package models

import "time"

type Ticket struct {
	ID           int64      `json:"id"`
	TicketNumber string     `json:"ticket_number"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`   // OPEN, ASSIGNED, CLOSED
	Priority     string     `json:"priority"` // LOW, MEDIUM, HIGH
	Category     string     `json:"category"` // AMC, PRODUCT, SOLUTION/SERVICE	
	CreatedBy    string     `json:"created_by"`
	AssignedTo   string     `json:"assigned_to"`
	CompanyID    int64      `json:"company_id"`
	ProductID    int64      `json:"product_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
}