package routes

import "github.com/gin-gonic/gin"

func RequestRoutes(server *gin.Engine) {
server.GET("/shows", getShows)
server.POST("/shows", createShows)
server.GET("/shows/:id", getShow)
}

