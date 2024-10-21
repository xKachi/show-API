package routes

import (
	"api3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "----"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "----"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "----"})

}