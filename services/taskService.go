package services

import (
	"ToDoList/models"
	"ToDoList/repositories"
)

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (t taskService) GetAll() ([]models.Task, error) {
	return t.repo.GetAll()
}
