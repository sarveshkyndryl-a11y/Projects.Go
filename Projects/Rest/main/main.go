package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type TeacherInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		w.Write([]byte("Hello root"))
		fmt.Println("Hello root triggered")
	})

	http.HandleFunc("/teachers/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:

			//query parameters
			///teachers/?key=value&key=value
			queryparams := r.URL.Query()
			search := queryparams.Get("search")
			fmt.Fprintln(w,search)
			fmt.Fprintln(w,r.URL.Path)
			path := strings.TrimPrefix(r.URL.Path,"/teachers/")
			id := strings.TrimSuffix(path,"/")
			fmt.Fprintln(w, "Teacher GET Method",id)

		case http.MethodPost:
			defer r.Body.Close()

			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			var teacher TeacherInfo
			err = json.Unmarshal(body, &teacher)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			fmt.Println("Received Teacher:", teacher)
			fmt.Fprintln(w, "Teacher POST Method")
			fmt.Fprintf(w, "Teacher Name: %s\n", teacher.Name)

		case http.MethodPatch:
			fmt.Fprintln(w, "Teacher PATCH Method")

		case http.MethodDelete:
			fmt.Fprintln(w, "Teacher DELETE Method")

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := 3000
	fmt.Println("Server started on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
