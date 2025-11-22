package main

import (
	"net/http"
	"taskmanager/internal/database"
	"taskmanager/internal/handlers"
)

func main() {
	storage := database.NewTaskStorage()
	taskHandlers := handlers.NewTaskHandlers(storage)

	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			taskHandlers.GetAllTasks(w, r)
		case "POST":
			taskHandlers.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/task", taskHandlers.GetTaskByID)

	http.HandleFunc("/", taskHandlers.ServeHTML)

	http.ListenAndServe(":8080", nil)
}
