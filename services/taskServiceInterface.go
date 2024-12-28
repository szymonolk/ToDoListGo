package services

import "ToDoList/models"

type TaskService interface {
	GetAll() ([]models.Task, error)
}
