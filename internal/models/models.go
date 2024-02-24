package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Импорт драйвера SQLite
)

// InitDB инициализирует подключение к базе данных
func InitDB(dataSourceName string) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
