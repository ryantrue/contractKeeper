package models

// User структура представляет пользователя в системе
type User struct {
	ID       int    `json:"id"`       // Уникальный идентификатор пользователя
	Username string `json:"username"` // Имя пользователя, используемое при входе
	Password string `json:"-"`        // Пароль пользователя, не включается в JSON представления
}
