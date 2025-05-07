package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB not responding:", err)
	}

	fmt.Println("✅ Connected to MongoDB")
	taskCollection = client.Database("gotasksdb").Collection("tasks")

	// Set up Gin
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from GoTasks"})
	})

	// Start the server
	router.Run(":8080")
}
