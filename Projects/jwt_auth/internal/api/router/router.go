package router

import (
	"jwtauth/internal/api/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/home", handlers.HomeHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	return mux
}
