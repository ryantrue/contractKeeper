package main

import (
	"log"
	"net/http"

	"github.com/RyanTrue/contractKeeper/internal/server/app"
)

func main() {
	application := app.NewApp() // Инициализация приложения

	log.Println("Запуск сервера на :8080")
	if err := http.ListenAndServe(":8080", application.Router); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s\n", err)
	}
}
