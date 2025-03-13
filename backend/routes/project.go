package main

import (
	"encoding/json"
	"net/http"
)

// Project represents a project in the system
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// In-memory project store (for demonstration purposes)
var projects = []Project{}

// GetProjectsHandler handles fetching all projects
func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

// CreateProjectHandler handles creating a new project
func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var newProject Project
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add project to the in-memory store
	projects = append(projects, newProject)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProject)
}
