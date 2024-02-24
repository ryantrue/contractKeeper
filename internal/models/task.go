package models

import "time"

// Task структура представляет задачу в системе
type Task struct {
	ID          int       `json:"id"`          // Уникальный идентификатор задачи
	Title       string    `json:"title"`       // Заголовок задачи
	Description string    `json:"description"` // Описание задачи
	IsCompleted bool      `json:"isCompleted"` // Статус выполнения задачи
	CreatedAt   time.Time `json:"createdAt"`   // Время создания задачи
	UpdatedAt   time.Time `json:"updatedAt"`   // Время последнего обновления задачи
}
