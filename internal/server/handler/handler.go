package handler

import (
	"net/http"
)

// HandleContracts выводит список контрактов.
// Эту функцию нужно будет зарегистрировать в роутере.
func HandleContracts(w http.ResponseWriter, r *http.Request) {
	// Здесь реализуйте логику получения и отправки контрактов
	w.Write([]byte("Список контрактов"))
}
