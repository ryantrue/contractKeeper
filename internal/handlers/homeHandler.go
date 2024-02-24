package handlers

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"html/template"
	"net/http"
	"path/filepath"
)

// Главная страница
func HomePageHandler(w http.ResponseWriter, r *http.Request, localizer *i18n.Localizer) {
	tmplPath := filepath.Join("templates", "login.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
