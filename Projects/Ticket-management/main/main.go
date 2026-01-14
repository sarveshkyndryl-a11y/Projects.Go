package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"ticket/internal/api/router"
	"ticket/internal/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	
	db.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	fmt.Println(" Ticket Management System")
	fmt.Println("Server running on port:", port)

	
	http.ListenAndServe(":"+port, router.Router())
}















