package main 

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
)

func main () {
	server := gin.Default()

	server.GET("/shows", getShows)

	server.Run(":8080")
}

func getShows(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Request Processed Successfully"})
}

