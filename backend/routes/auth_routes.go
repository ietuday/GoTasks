package routes

import (
	"gotasks/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, userCollection *mongo.Collection) {
	h := handlers.NewAuthHandler(userCollection)

	rg.POST("/register", h.Register)
	rg.POST("/login", h.Login)
}
