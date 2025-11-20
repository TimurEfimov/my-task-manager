package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/internal/database"
)

type TaskHandlers struct {
	storage *database.TaskStorage
}

func NewTaskHandlers(storage *database.TaskStorage) *TaskHandlers {
	return &TaskHandlers{
		storage: storage,
	}
}

func (h *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := h.storage.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}