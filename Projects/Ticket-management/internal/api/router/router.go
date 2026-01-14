package router

import (
	"net/http"
	"ticket/internal/api/handlers"
)

func Router()*http.ServeMux{
mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/tickets/", handlers.TicketHandler)

	
	// companies Collection routes
	mux.HandleFunc("GET /companies", handlers.GetAllCompaniesHandler)
	mux.HandleFunc("POST /companies", handlers.CreateCompanyHandler)

	mux.HandleFunc("PUT /companies/{id}", handlers.UpdateCompanyHandler)
	mux.HandleFunc("PATCH /companies/{id}", handlers.PatchCompanyHandler)
	mux.HandleFunc("DELETE /companies/{id}", handlers.DeleteCompanyHandler)


	return mux
}