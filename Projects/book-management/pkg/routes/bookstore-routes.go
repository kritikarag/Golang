package routes

import (
	"github.com/gin-gonic/gin"
	"../controllers"
	"github.com/kritikarag/golang/Projects/book-management/pkg/controllers"
)

func RegisterBookStoreRoutes(c *gin.Context){
	r:=gin.Default()

	r.POST("/book",controllers.CreateBook)
	r.GET("/book",controllers.GetBook)
	r.GET("/book/:id",controllers.GetBookById)
	r.PUT("/book/:id",controllers.GetUpdateBook)
	r.DELETE("/book/:id",controllers.DeleteBook)
}