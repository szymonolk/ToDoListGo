package repositories

import "ToDoList/models"

type TaskRepository interface {
	GetAll() ([]models.Task, error)
}
