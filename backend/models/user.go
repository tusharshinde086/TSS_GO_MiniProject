package models

// User represents a user in the system
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"` // In a real application, do not store passwords in plain text
}