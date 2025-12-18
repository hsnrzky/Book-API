package routes

import (
	"github.com/gin-gonic/gin"
	"golang-api/controllers"
)

func BookRoutes(r *gin.Engine) {
	r.POST("/books", controllers.AddBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
