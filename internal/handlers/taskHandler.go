// internal/handlers/taskHandler.go

package handlers

import (
	"encoding/json"
	"net/http"
)

// Task структура для представления задачи
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Другие поля...
}

// GetTasksHandler возвращает список задач
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь должна быть логика получения задач из базы данных...

	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1"},
		// Другие задачи...
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
