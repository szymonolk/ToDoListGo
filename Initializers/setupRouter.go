package Initializers

import (
	"ToDoList/controllers"
	"ToDoList/repositories"
	"ToDoList/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	taskRepository := repositories.NewTaskRepository(DB)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	r.GET("/task/all", taskController.GetAllTasks)
}
