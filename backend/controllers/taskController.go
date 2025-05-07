package controllers

import (
	"gotasks/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{Title: "Learn Go", Description: "Go is awesome", Completed: false},
	{Title: "Learn Docker", Description: "Containers are life", Completed: false},
}

// GetTasks returns a list of tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// AddTask adds a new task to the list
func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}
