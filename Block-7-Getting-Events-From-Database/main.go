package main

import (
	"api3/db"
	"api3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main () {
	db.InitDB()
	server := gin.Default()

	server.GET("/shows", getShows)
	server.GET("/shows", createShows)

	server.Run(":8080")
}

func getShows(context *gin.Context) {
	shows, err := models.GetAllShows()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the shows"})
		return
	}

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

	err = show.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to saved to the database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Show created!", "Show": show})
}

