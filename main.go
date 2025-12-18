package main

import (
	"github.com/gin-gonic/gin"
	"golang-api/config"
	"golang-api/models"
	"golang-api/routes"
)

func main() {
	// koneksi database
	config.ConnectDatabase()

	// auto migrate
	config.DB.AutoMigrate(&models.Book{})

	// gin
	r := gin.Default()
	routes.BookRoutes(r)

	r.Run(":8080")
}
