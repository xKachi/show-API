package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/shows", getShows)
	server.GET("/shows/:id", getShow)
	server.POST("/shows", createShow)
	server.PUT("/shows/:id", updateShow)
	server.DELETE("/shows/:id", deleteShow)

	// User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}

