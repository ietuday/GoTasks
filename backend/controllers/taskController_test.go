package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gotasks/models"

	"github.com/gin-gonic/gin"
)

// ===== Mock Collection =====

// Mock collection simulates the MongoDB collection operations
type mockCollection struct {
	insertFunc  func(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	findFunc    func(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
	deleteFunc  func(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	findOneFunc func(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
	updateFunc  func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

// ===== Mock Mongo Cursor Wrapper =====

// mockMongoCursor wraps mockCursor to implement mongo.Cursor interface
type mockMongoCursor struct {
	*mockCursor
}

// Implement the required methods of mongo.Cursor
func (m *mockMongoCursor) ID() int64 {
	return 0
}

func (m *mockMongoCursor) Current() bson.Raw {
	return nil
}

func (m *mockMongoCursor) Err() error {
	return nil
}

func (m *mockMongoCursor) Next(ctx context.Context) bool {
	return m.mockCursor.Next(ctx)
}

func (m *mockMongoCursor) Decode(val interface{}) error {
	return m.mockCursor.Decode(val)
}

func (m *mockMongoCursor) Close(ctx context.Context) error {
	return m.mockCursor.Close(ctx)
}

// Mock UpdateOne method
func (m *mockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, filter, update, opts...)
	}
	return &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil
}

// Mock FindOne method
func (m *mockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	if m.findOneFunc != nil {
		return m.findOneFunc(ctx, filter, opts...)
	}
	return nil
}

// Mock DeleteOne method
func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, filter, opts...)
	}
	return &mongo.DeleteResult{DeletedCount: 1}, nil
}

// Mock InsertOne method
func (m *mockCollection) InsertOne(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	if m.insertFunc != nil {
		return m.insertFunc(ctx, doc, opts...)
	}
	return &mongo.InsertOneResult{}, nil
}

// Mock Find method
func (m *mockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	if m.findFunc != nil {
		return m.findFunc(ctx, filter, opts...)
	}
	return nil, nil
}

// ===== Mock Cursor =====

// Mock cursor simulates the behavior of a MongoDB cursor.
type mockCursor struct {
	data []models.Task
	pos  int
}

// Simulate moving to the next item in the cursor
func (c *mockCursor) Next(_ context.Context) bool {
	c.pos++
	return c.pos <= len(c.data)
}

// Simulate decoding the current cursor item into a Task object
func (c *mockCursor) Decode(val interface{}) error {
	if c.pos-1 >= len(c.data) {
		return errors.New("out of bounds")
	}
	taskPtr, ok := val.(*models.Task)
	if !ok {
		return errors.New("invalid decode target")
	}
	*taskPtr = c.data[c.pos-1]
	return nil
}

// Simulate closing the cursor
func (c *mockCursor) Close(_ context.Context) error {
	return nil
}

// ===== Mock SingleResult =====

type mockSingleResult struct {
	decodeFunc func(v interface{}) error
}

// Decode method for mockSingleResult
func (m *mockSingleResult) Decode(v interface{}) error {
	if m.decodeFunc == nil {
		return errors.New("decodeFunc is not initialized")
	}
	return m.decodeFunc(v)
}

// ======= TEST: GetTasks =======

// Test for the GetTasks endpoint
func TestGetTasks(t *testing.T) {
	mockData := []models.Task{
		{Title: "Mock Task 1", Description: "Testing 123"},
		{Title: "Mock Task 2", Description: "Another test"},
	}

	mockCol := &mockCollection{
		findFunc: func(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
			cursor := &mockCursor{data: mockData}

			// Convert []models.Task to []interface{}
			documents := make([]interface{}, len(cursor.data))
			for i, task := range cursor.data {
				documents[i] = task
			}

			return mongo.NewCursorFromDocuments(documents, nil, nil)
		},
	}

	InitController(mockCol)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks", nil) // Ensure the request method and path are correct
	GetTasks(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Parse the response body
	var tasks []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	// Validate the response data
	if len(tasks) != len(mockData) {
		t.Fatalf("expected %d tasks, got %d", len(mockData), len(tasks))
	}

	for i, task := range tasks {
		if task.Title != mockData[i].Title || task.Description != mockData[i].Description {
			t.Errorf("unexpected task at index %d: got %+v, want %+v", i, task, mockData[i])
		}
	}
}

// ======= TEST: AddTask =======

// Test for the AddTask endpoint
func TestAddTask(t *testing.T) {
	mockCol := &mockCollection{
		insertFunc: func(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return &mongo.InsertOneResult{}, nil
		},
	}
	InitController(mockCol)

	task := models.Task{
		Title:       "Unit Test Task",
		Description: "AddTask test case",
	}
	body, _ := json.Marshal(task)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	AddTask(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", w.Code)
	}

	var result models.Task
	err := json.NewDecoder(io.NopCloser(w.Body)).Decode(&result)
	if err != nil || result.Title != task.Title {
		t.Fatalf("unexpected response body: %s", w.Body.String())
	}
}

// ======= TEST: EditTask =======

// Test for the EditTask endpoint
func TestEditTask(t *testing.T) {
	originalID := primitive.NewObjectID()

	mockCol := &mockCollection{
		updateFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
			// ✅ Simulate a successful update
			return &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}, nil
		},
		findOneFunc: func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
			// ✅ Simulate the updated task being returned
			mockDoc := bson.M{
				"_id":         originalID,
				"title":       "Updated Task",
				"description": "Updated description",
			}
			return mongo.NewSingleResultFromDocument(mockDoc, nil, nil)
		},
	}
	InitController(mockCol)

	// Prepare request body
	updatedTask := models.Task{
		Title:       "Updated Task",
		Description: "Updated description",
	}
	body, _ := json.Marshal(updatedTask)

	// Gin test setup
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("PUT", "/tasks/"+originalID.Hex(), bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	// 👇 This is crucial
	c.Params = gin.Params{gin.Param{Key: "id", Value: originalID.Hex()}}

	// Act
	EditTask(c)

	// Assert
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var result models.Task
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result.Title != updatedTask.Title || result.Description != updatedTask.Description {
		t.Errorf("unexpected task: got %+v, want %+v", result, updatedTask)
	}
}

// ======= TEST: GetTaskDetail =======

// Test for the GetTaskDetail endpoint
func TestGetTaskDetail(t *testing.T) {
	// Setup
	objectID := primitive.NewObjectID()
	mockData := models.Task{
		ID:          objectID,
		Title:       "Task 1",
		Description: "Detail of Task 1",
	}

	mockCol := &mockCollection{
		findOneFunc: func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
			// Validate filter if you want:
			expectedFilter := bson.D{{Key: "_id", Value: objectID}}
			if !reflect.DeepEqual(filter, expectedFilter) {
				t.Errorf("unexpected filter: got %v, want %v", filter, expectedFilter)
			}

			// Mocked behavior
			mockDoc := bson.M{
				"_id":         mockData.ID,
				"title":       mockData.Title,
				"description": mockData.Description,
			}
			return mongo.NewSingleResultFromDocument(mockDoc, nil, nil)
		},
	}

	InitController(mockCol)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/tasks/"+objectID.Hex(), nil)
	// This is crucial for extracting :id
	c.Params = gin.Params{gin.Param{Key: "id", Value: objectID.Hex()}}

	// Call the controller
	GetTaskDetail(c)

	// Verify response
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var result models.Task
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if result.ID != mockData.ID || result.Title != mockData.Title || result.Description != mockData.Description {
		t.Fatalf("unexpected response body: got %+v, want %+v", result, mockData)
	}
}

// ======= TEST: DeleteTask =======

// Test for the DeleteTask endpoint
// ======= TEST: DeleteTask =======

// Test for the DeleteTask endpoint
func TestDeleteTask(t *testing.T) {
	validID := "507f1f77bcf86cd799439011" // A valid ObjectID
	objectID, _ := primitive.ObjectIDFromHex(validID)

	mockCol := &mockCollection{
		deleteFunc: func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
			expectedFilter := bson.D{{Key: "_id", Value: objectID}} // Match the type used in real code

			// Convert both to bson.D to compare properly
			actualFilter, ok := filter.(bson.D)
			if !ok {
				t.Errorf("unexpected filter type: got %T, want bson.D", filter)
			}

			if !reflect.DeepEqual(actualFilter, expectedFilter) {
				t.Errorf("unexpected filter: got %v, want %v", actualFilter, expectedFilter)
			}

			return &mongo.DeleteResult{DeletedCount: 1}, nil
		},
	}
	InitController(mockCol)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("DELETE", "/tasks/"+validID, nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: validID}}

	DeleteTask(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Code)
	}

	expected := `{"message":"Task deleted successfully"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("unexpected response body: got %s, want %s", w.Body.String(), expected)
	}
}
