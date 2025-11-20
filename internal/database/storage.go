package database

import (
	"sync"
	"taskmanager/internal/models"
	"time"
)

type TaskStorage struct {
	tasks  map[int]models.Task
	nextID int
	mu     sync.RWMutex
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *TaskStorage) CreateTask(title, description string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := models.Task{
		ID:        s.nextID,
		Title:     title,
		Desc:      description,
		Status:    models.Todo,
		CreatedAt: time.Now(),
	}

	s.tasks[s.nextID] = task

	s.nextID++

	return task
}

func (s *TaskStorage) GetAllTasks() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskStorage) GetTaskByID(id int) *models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	
	if !exists {
		return nil
	}

	return &task
}
