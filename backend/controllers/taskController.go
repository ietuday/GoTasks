package controllers

import (
	"context"
	"net/http"

	"gotasks/models" // Importing the Task model which defines task data

	"github.com/gin-gonic/gin"                  // Web framework for building RESTful APIs
	"go.mongodb.org/mongo-driver/bson"          // MongoDB BSON helpers for structuring queries
	"go.mongodb.org/mongo-driver/mongo"         // MongoDB driver for Go
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDB connection and query options
)

// ✅ Interface with correct variadic method signatures
// Define an interface that describes the methods available on our Task collection.
// This allows us to decouple the controller from the database, making it testable and flexible.
type TaskCollection interface {
	InsertOne(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

// Global variable to hold the injected collection object
// This is set in the InitController function when we call it from main.go.
var taskCol TaskCollection

// ====================
// Controller Initialization
// ====================

// InitController is called from main.go to inject the MongoDB collection into the controller
// This function sets the taskCol variable with the actual collection provided by main.go
func InitController(col TaskCollection) {
	taskCol = col
}

// ====================
// 🚀 GetTasks Endpoint
// ====================

// GetTasks retrieves all tasks from the MongoDB collection and sends them in the response.
// This endpoint is exposed as a GET route that returns a list of all tasks stored in the database.
func GetTasks(c *gin.Context) {
	// Fetch all tasks from the MongoDB collection using an empty filter (bson.D{} means "all records")
	cursor, err := taskCol.Find(context.Background(), bson.D{})
	if err != nil {
		// If an error occurs while fetching tasks, return a 500 Internal Server Error response.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks: " + err.Error()})
		return
	}

	// Declare a slice to hold the tasks that will be retrieved from the database
	var tasks []models.Task

	// Parse the results from the cursor into the tasks slice
	if err := cursor.All(context.Background(), &tasks); err != nil {
		// If an error occurs during parsing, return a 500 Internal Server Error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks: " + err.Error()})
		return
	}

	// Successfully retrieved and parsed tasks. Return them with a 200 OK status.
	c.JSON(http.StatusOK, tasks)
}

// ====================
// ➕ AddTask Endpoint
// ====================

// AddTask allows a client to add a new task to the MongoDB collection.
// It is exposed as a POST route, expecting a JSON payload that represents the new task.
func AddTask(c *gin.Context) {
	// Create an empty Task object to bind the incoming JSON data to
	var newTask models.Task

	// Bind the incoming JSON body to the newTask struct
	// If binding fails (invalid JSON format or missing required fields), return a 400 Bad Request response
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Insert the new task into the MongoDB collection
	_, err := taskCol.InsertOne(context.Background(), newTask)
	if err != nil {
		// If insertion fails, return a 500 Internal Server Error with a detailed error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert task: " + err.Error()})
		return
	}

	// Successfully added the task, return it with a 201 Created status
	// This indicates that the task has been successfully created and stored in the database
	c.JSON(http.StatusCreated, newTask)
}
