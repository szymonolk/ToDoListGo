package repositories

import (
	"ToDoList/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (t taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := t.db.Find(&tasks).Error
	return tasks, err
}
