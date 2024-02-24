package data

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string // Хранит хешированный пароль
}

// GetUserByUsername возвращает пользователя по его имени пользователя, если он существует
func GetUserByUsername(username string) (*User, error) {
	// заглушка для демонстрации
	return &User{
		ID:       1,
		Username: "testUser",
		Password: "$2a$10$E8jg/JPOM31Ak.hFfbRbfeJ8z1w5XhiF9JkKthB3.lRyMXqBC/7O6", // Пример хешированного пароля
	}, nil
}

// ValidateUserCredentials проверяет учетные данные пользователя
func ValidateUserCredentials(username, password string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}

	// Сравниваем хешированный пароль с предоставленным паролем
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// AddUser добавляет нового пользователя в базу данных
func AddUser(username, hashedPassword string) error {
	// Здесь должна быть логика для добавления пользователя в базу данных
	// Возвращаем nil для демонстрации, предполагая, что пользователь успешно добавлен
	return nil
}
