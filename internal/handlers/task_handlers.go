package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"taskmanager/internal/database"
)

type TaskHandlers struct {
	storage *database.TaskStorage
}

type CreateTaskRequest struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func NewTaskHandlers(storage *database.TaskStorage) *TaskHandlers {
	return &TaskHandlers{
		storage: storage,
	}
}

func (h *TaskHandlers) ServeHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func (h *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := h.storage.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task := h.storage.CreateTask(req.Title, req.Desc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	task := h.storage.GetTaskByID(id)
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}