package controllers

import (
	"task-management/models"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"gin.com/gin/models"
)

type CreateTaskInput struct {
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
}

type UpdateTaskInput struct {
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
}

func GetTasks(c *gin.Context){
	var tasks []models.Task
	//db := c.MustGet("db").(*gorm.DB)
	models.DB.Find(&tasks)
	c.JSON(http.StatusOK,gin.H{"data": tasks})

}

func CreateTask(c *gin.Context){
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currTime := time.Now();
	task:= models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail, Date: currTime.Format("2006-01-02 15:04:05")}
	models.DB.Create(&task)
}

func GetTaskById(c *gin.Context){
	var task models.Task
	//id := c.Request.URL.Query().Get("id")
	db := c.MustGet("db").(*gorm.DB)
	//db := c.MustGet("db").(*gorm.DB)
	if err:= db.Where("id = ?",c.Param("id")).First(&task).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Record not found!"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"data":task})
}

func UpdateTask(c *gin.Context){

	var task models.Task
	
	id:=c.Request.URL.Query().Get("id")

	if err:= models.DB.Where("id = ?",id).First(&task).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Record not found!"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTask models.Task

	updatedTask.TaskName = input.TaskName
	updatedTask.TaskDetail = input.TaskDetail
	updatedTask.Date = time.Now().Format("2006-01-02 15:04:05")

	models.DB.Model(&task).Updates(updatedTask)
	c.JSON(http.StatusOK,gin.H{"data":task})
}

func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}