package models

// Task represents a task in a project
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ProjectID   int    `json:"project_id"` // Foreign key to associate the task with a project
}
