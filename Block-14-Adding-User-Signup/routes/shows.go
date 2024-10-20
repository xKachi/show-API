package routes

import (
	"api3/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func getShows(context *gin.Context) {
	shows, err := models.GetAllShows()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get the shows"})
		return
	}

	context.JSON(http.StatusOK, shows)
}

func getShow(context *gin.Context) {
	showid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request id"})
		return
	}

	show, err := models.GetShowID(showid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the show with this id"})
		return
	}

	context.JSON(http.StatusOK, show)
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

func updateShow(context *gin.Context) {
	showid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request id"})
		return
	}

	_, err = models.GetShowID(showid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the show with this id"})
		return
	}

	var updateShow models.Show

	context.ShouldBindJSON(&updateShow)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get the show with this id"})
		return
	}

	updateShow.ID = showid
	err = updateShow.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "The show could not be updated"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Show has been updated successfully"})
}

func deleteShow(context *gin.Context) {
	showid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request id"})
		return
	}

	show, err := models.GetShowID(showid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the show with this id"})
		return
	}

	err = show.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Show deletion failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "show deletion successful"})

}
