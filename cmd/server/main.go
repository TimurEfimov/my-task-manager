package main

import (
	"net/http"
	"taskmanager/internal/database"
	"taskmanager/internal/handlers"
)

func main() {
	storage := database.NewTaskStorage()

	taskHandlers := handlers.NewTaskHandlers(storage)

	http.HandleFunc("/tasks", taskHandlers.GetAllTasks)

	http.ListenAndServe(":8080", nil)
}
