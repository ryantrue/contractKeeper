package app

import (
	"github.com/gorilla/mux" // Используем Gorilla Mux для маршрутизации
)

// App содержит компоненты вашего приложения.
type App struct {
	Router *mux.Router
}

// NewApp инициализирует новое приложение с необходимыми компонентами.
func NewApp() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.setRoutes()
	return a
}

// setRoutes определяет маршруты приложения.
func (a *App) setRoutes() {
	// Здесь добавьте маршруты, например:
	// a.Router.HandleFunc("/contracts", a.handleContracts).Methods("GET")
}

// Здесь добавьте обработчики для ваших эндпоинтов, например:
// func (a *App) handleContracts(w http.ResponseWriter, r *http.Request) {
//     // Логика обработки запросов
// }
