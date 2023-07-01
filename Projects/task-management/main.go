package main

import (
	"fmt"
	"log"
	"task-management/controllers"
	"task-management/models"

	"github.com/gin-gonic/gin"
)

func main() { 
	db, err := models.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&models.Task{})
	
	r := gin.Default()
	// r.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// })

	r.GET("/tasks",controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.POST("/tasks", controllers.CreateTask) 
	r.PUT("/tasks/:id", controllers.UpdateTask) 
	r.DELETE("/tasks/:id",controllers.DeleteTask)

	fmt.Println("Starting Server at port no. 8084")

	r.Run(":8084")
}