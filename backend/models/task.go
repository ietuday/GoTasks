package models

import "strings"

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Basic validation
func (t *Task) Validate() bool {
	return strings.TrimSpace(t.Title) != ""
}
