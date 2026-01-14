package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"Restapi/internal/api/middlewares"
)

/* =======================
   MODELS
======================= */

type Teacher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Class   string `json:"class"`
}

/* =======================
   IN-MEMORY STORE
======================= */

var teachers = make(map[int]Teacher)
var nextID = 1

func init() {
	teachers[nextID] = Teacher{
		ID:      nextID,
		Name:    "John Doe",
		Subject: "Mathematics",
		Class:   "10A",
	}
	nextID++

	teachers[nextID] = Teacher{
		ID:      nextID,
		Name:    "Jane Smith",
		Subject: "Science",
		Class:   "10B",
	}
	nextID++
}

/* =======================
   HANDLERS
======================= */

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Optional query param
	class := r.URL.Query().Get("class")

	// Extract ID from path: /teachers/{id}
	path := strings.TrimPrefix(r.URL.Path, "/teachers")
	path = strings.Trim(path, "/")

	// 🔹 CASE 1: GET /teachers/{id}
	if path != "" {
		id, err := strconv.Atoi(path)
		if err != nil {
			http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
			return
		}

		teacher, exists := teachers[id]
		if !exists {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(teacher)
		return
	}

	// 🔹 CASE 2: GET /teachers OR /teachers?class=10A
	teachersList := make([]Teacher, 0)

	for _, teacher := range teachers {
		if class == "" || teacher.Class == class {
			teachersList = append(teachersList, teacher)
		}
	}

	response := struct {
		Status   string    `json:"status"`
		Count    int       `json:"count"`
		Teachers []Teacher `json:"teachers"`
	}{
		Status:   "success",
		Count:    len(teachersList),
		Teachers: teachersList,
	}

	json.NewEncoder(w).Encode(response)
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getTeachersHandler(w, r)

	case http.MethodPost:
		http.Error(w, "POST not implemented", http.StatusNotImplemented)

	case http.MethodPatch:
		http.Error(w, "PATCH not implemented", http.StatusNotImplemented)

	case http.MethodDelete:
		http.Error(w, "DELETE not implemented", http.StatusNotImplemented)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root"))
}

/* =======================
   MAIN
======================= */

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers", teachersHandler)
	mux.HandleFunc("/teachers/", teachersHandler)

	port := 3000
	fmt.Println("Server started on port", port)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		middlewares.Cors(mux),
	)

	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
