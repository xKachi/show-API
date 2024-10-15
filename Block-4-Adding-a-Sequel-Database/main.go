package main

import (
	"api3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main () {
	server := gin.Default()

	server.GET("/shows", getShows)
	server.GET("/shows", createShows)

	server.Run(":8080")
}

func getShows(context *gin.Context) {
	shows := models.GetAllShows()
	context.JSON(http.StatusOK, shows)
}

func createShows(context *gin.Context) {
	var show models.Show

	err := context.ShouldBindJSON(&show)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse request"})
		return
	}

	show.ID = 1
	show.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Show created!", "Show": show})
}

