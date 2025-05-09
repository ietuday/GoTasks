package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gotasks/controllers" // Add to imports

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"                  // Web framework for building APIs
	"go.mongodb.org/mongo-driver/mongo"         // MongoDB driver
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDB connection options
)

// Global variable to hold a reference to the MongoDB "tasks" collection
var taskCollection *mongo.Collection

func main() {
	// =======================
	// 🔌 Connect to MongoDB
	// =======================

	// Set the MongoDB URI — "mongo" is the service name (e.g., in Docker)
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	// Create a context with a 10-second timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Make sure the context is cleaned up

	// Try to connect to MongoDB using the URI and context
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err) // Exit if connection fails
	}

	// Ping the database to make sure it's alive and responsive
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB not responding:", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	// Access the database "gotasksdb" and the collection "tasks"
	taskCollection = client.Database("gotasksdb").Collection("tasks")

	// ========================
	// 🌐 Set up Gin Web Server
	// ========================

	// Initialize the default Gin router (includes logger and recovery middleware)
	router := gin.Default()

	// 💥 CORS middleware here
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint — hit this to verify the server is running
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from GoTasks"})
	})

	// Pass collection to controller
	controllers.InitController(taskCollection)

	// Define routes
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.AddTask)
	router.PUT("/tasks/:id", controllers.EditTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.GET("/tasks/:id", controllers.GetTaskDetail)

	// ========================
	// 🚀 Start HTTP Server
	// ========================

	// Run the server on port 8080
	router.Run(":8080")
}
