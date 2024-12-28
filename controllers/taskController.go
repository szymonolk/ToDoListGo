package controllers

import (
	"ToDoList/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{service}
}

func (controller *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
