package routes

import (
	"encoding/json"
	"net/http"
	"your-module-name/models" // Adjust the import path as necessary
	"your-module-name/utils"  // Adjust the import path as necessary
)

// In-memory user store (for demonstration purposes)
var users = []models.User{}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	newUser.Password = hashedPassword

	// Add user to the in-memory store
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user exists
	for _, user := range users {
		if user.Email == loginUser.Email && utils.CheckPasswordHash(loginUser.Password, user.Password) {
			// Here you would typically generate a token and return it
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user) // Return user info (excluding password)
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
