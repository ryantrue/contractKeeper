package data

import "github.com/ryantrue/contractKeeper.git/internal/models"

type Task struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
}

// GetTasks возвращает список всех задач
func GetTasks() ([]models.Task, error) {
	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", IsCompleted: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", IsCompleted: true},
	}
	return tasks, nil
}

func AddTask(task models.Task) error {
	// Пример реализации. В реальном приложении здесь будет запрос к базе данных.
	return nil // Предполагается, что задача успешно добавлена
}

func DeleteTask(taskID int) error {
	// Пример реализации. В реальном приложении здесь будет запрос к базе данных.
	return nil // Предполагается, что задача успешно удалена
}
