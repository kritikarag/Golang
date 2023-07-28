package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data":"HEllo Everyone!"})
	})

	server.GET("/books")
	server.Run(":8008")
}