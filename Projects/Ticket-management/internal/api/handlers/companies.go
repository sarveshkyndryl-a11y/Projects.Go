package handlers

import (
	"encoding/json"
	"net/http"

	"ticket/internal/models"
	"ticket/internal/repo"

	"github.com/google/uuid"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if company.Name == "" || company.Contact_email== "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	if err := repo.CreateCompany(r.Context(), &company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func GetAllCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	companies, err := repo.GetAllCompanies(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(companies)
}

func UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid company ID format", http.StatusBadRequest)
		return
	}

	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	company.ID = uid

	if err := repo.UpdateCompany(r.Context(), &company); err != nil {
		http.Error(w, "Failed to update company", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(company)
}

func PatchCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	var updates map[string]any
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := repo.PatchCompany(r.Context(), id, updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Company updated",
	})
}

func DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	if err := repo.DeleteCompany(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
