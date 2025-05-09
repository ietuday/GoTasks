package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gotasks/models"

	"github.com/gin-gonic/gin"
)

// ===== Mock Collection =====

// Mock collection simulates the MongoDB collection operations
type mockCollection struct {
	// Function to mock InsertOne behavior
	insertFunc func(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	// Function to mock Find behavior
	findFunc func(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
}

// Mock InsertOne method to simulate MongoDB's InsertOne operation
func (m *mockCollection) InsertOne(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	// If insertFunc is defined, call it, otherwise return default InsertOneResult
	if m.insertFunc != nil {
		return m.insertFunc(ctx, doc, opts...)
	}
	return &mongo.InsertOneResult{}, nil
}

// Mock Find method to simulate MongoDB's Find operation
func (m *mockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	// If findFunc is defined, call it, otherwise return nil (no cursor)
	if m.findFunc != nil {
		// We can return a cursor from mock data or simulate an error
		return m.findFunc(ctx, filter, opts...)
	}
	return nil, nil
}

// ===== Mock Cursor =====

// Mock cursor simulates the behavior of a MongoDB cursor.
type mockCursor struct {
	data []models.Task // Mock data that the cursor will iterate over
	pos  int           // Position of the current item in the mock data
}

// Simulate moving to the next item in the cursor
func (c *mockCursor) Next(_ context.Context) bool {
	c.pos++ // Move to the next task in the list
	// If the current position is within bounds, return true
	return c.pos <= len(c.data)
}

// Simulate decoding the current cursor item into a Task object
func (c *mockCursor) Decode(val interface{}) error {
	// If the cursor position exceeds the data length, return an out-of-bounds error
	if c.pos-1 >= len(c.data) {
		return errors.New("out of bounds")
	}
	// Ensure the target value is a pointer to a Task
	taskPtr, ok := val.(*models.Task)
	if !ok {
		return errors.New("invalid decode target")
	}
	// Assign the task data from the mock data slice
	*taskPtr = c.data[c.pos-1]
	return nil
}

// Simulate closing the cursor (no-op here)
func (c *mockCursor) Close(_ context.Context) error {
	return nil
}

// ======= TEST: GetTasks =======

// Test for the GetTasks endpoint, which retrieves all tasks from the collection
func TestGetTasks(t *testing.T) {
	// Sample mock task data that represents what might be in the MongoDB collection
	mockData := []models.Task{
		{Title: "Mock Task 1", Description: "Testing 123"},
		{Title: "Mock Task 2", Description: "Another test"},
	}

	// Create a mock collection where the Find function returns mock data
	mockCol := &mockCollection{
		findFunc: func(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
			// Convert the mock task data into interface{} so that it can be used by the cursor
			interfaceData := make([]interface{}, len(mockData))
			for i, v := range mockData {
				interfaceData[i] = v
			}
			// Create and return a mock cursor using the mock data
			cursor, err := mongo.NewCursorFromDocuments(interfaceData, nil, nil)
			return cursor, err
		},
	}

	// Initialize the controller with the mock collection
	InitController(mockCol)

	// Create a new recorder to capture the HTTP response
	w := httptest.NewRecorder()
	// Create a Gin test context for handling requests and responses
	c, _ := gin.CreateTestContext(w)

	// Call the GetTasks function which is mapped to the GET /tasks route
	GetTasks(c)

	// Assert that the response status code is 200 OK
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Parse the JSON response body into a list of Task objects
	var tasks []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	// Check that the parsed tasks are correct and match the mock data
	if err != nil || len(tasks) != 2 {
		t.Fatalf("unexpected response: %v", w.Body.String())
	}
}

// ======= TEST: AddTask =======

// Test for the AddTask endpoint, which creates a new task
func TestAddTask(t *testing.T) {
	// Mock collection with an insert function that simulates successful insertion
	mockCol := &mockCollection{
		insertFunc: func(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			// Simulate successful insertion, returning a mock InsertOneResult
			return &mongo.InsertOneResult{}, nil
		},
	}
	// Initialize the controller with the mock collection
	InitController(mockCol)

	// Create a new task object to be added
	task := models.Task{
		Title:       "Unit Test Task",
		Description: "AddTask test case",
	}
	// Marshal the task object into a JSON request body
	body, _ := json.Marshal(task)

	// Create a recorder to capture the HTTP response
	w := httptest.NewRecorder()
	// Create a test context for the Gin request
	c, _ := gin.CreateTestContext(w)
	// Simulate a POST request to the /tasks endpoint with the task body
	c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	// Set the Content-Type header to application/json
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the AddTask function to handle the request
	AddTask(c)

	// Assert that the response status code is 201 Created
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", w.Code)
	}

	// Parse the response body to ensure the task returned matches the one created
	var result models.Task
	err := json.NewDecoder(io.NopCloser(w.Body)).Decode(&result)
	if err != nil || result.Title != task.Title {
		t.Fatalf("unexpected response body: %s", w.Body.String())
	}
}
