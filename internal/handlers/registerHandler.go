package handlers

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"html/template"
	"net/http"
	"path/filepath"
)

// Страница регистрации
func RegisterHandler(w http.ResponseWriter, r *http.Request, localizer *i18n.Localizer) {
	if r.Method == http.MethodPost {
		// Обработка данных формы регистрации
		// логика регистрации нового пользователя

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmplPath := filepath.Join("templates", "register.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
