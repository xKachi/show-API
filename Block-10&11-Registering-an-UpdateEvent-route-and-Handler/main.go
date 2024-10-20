package main

import (
	"api3/db"
	"api3/routes"

	"github.com/gin-gonic/gin"
)


func main () {
	db.InitDB()
	server := gin.Default()
	routes.RequestRoutes(server)

	server.Run(":8080")
}

