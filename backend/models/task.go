package models

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task struct defines the structure of a task object
// It will be used for the task data sent and received via the API.
type Task struct {
	// ID is a unique identifier for the task, mapped to MongoDB's _id field
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// Title is a string that represents the task's name or title
	Title string `json:"title"` // JSON tag maps the struct field to the corresponding JSON key
	// Description is a string that provides details about the task
	Description string `json:"description"` // This field is also mapped to the JSON key "description"
	// Completed is a boolean indicating whether the task has been completed or not
	Completed bool `json:"completed"` // Maps to the JSON key "completed"
}

// Validate method checks if the Title field is not empty or just spaces
// Returns true if the Title is valid (not empty or whitespace), false otherwise.
func (t *Task) Validate() bool {
	// The Validate function trims any leading or trailing spaces from the Title
	// and checks if the Title has any remaining non-space characters
	// If the Title is empty or just spaces, the task is considered invalid
	return strings.TrimSpace(t.Title) != ""
}
