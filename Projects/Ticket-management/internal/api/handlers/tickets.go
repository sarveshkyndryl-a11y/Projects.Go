package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"ticket/internal/models"

	"time"
)

var tickets = make(map[int64]models.Ticket)
var nextID int64 = 1
var mutex = &sync.Mutex{}


func TicketHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		GetAllTicketHandler(w, r)

	case http.MethodPost:
		AddTicketHandler(w, r)

	}
}





func GetAllTicketHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var path = strings.TrimPrefix(r.URL.Path, "/tickets/")
	path = strings.Trim(path, "/")
	var ticketList = make([]models.Ticket, 0, len(tickets))

    if path != "" {

		id, err := strconv.Atoi(path)
		if err != nil {
			http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
			return
		}
		ticket, exists := tickets[int64(id)]
		if !exists {
			http.Error(w, "Ticket not found", http.StatusNotFound)
			return
		}
		ticketList = append(ticketList, ticket)
		json.NewEncoder(w).Encode(ticketList)
		return
	} 
		
		for _, ticket := range tickets {
			ticketList = append(ticketList, ticket)
		}
		response:= struct{
			Status string
			Tickets []models.Ticket
		}{
			Status:  "success",
			Tickets: ticketList,
		}
		json.NewEncoder(w).Encode(response)
}

func AddTicketHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTicket models.Ticket

	if err := json.NewDecoder(r.Body).Decode(&newTicket); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	if newTicket.Title == "" || newTicket.Priority == "" || newTicket.Category == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
    newTicket.Status = "OPEN"
	
	newTicket.ID = nextID
	newTicket.TicketNumber = fmt.Sprintf("TICKET-%03d", nextID)
	newTicket.Status = "OPEN"
	newTicket.CreatedAt = time.Now()
	newTicket.UpdatedAt = time.Now()


	tickets[newTicket.ID] = newTicket
	nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string        `json:"status"`
		Ticket models.Ticket `json:"ticket"`
	}{
		Status: "success",
		Ticket: newTicket,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
