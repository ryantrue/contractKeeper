package handlers

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ryantrue/contractKeeper.git/internal/auth"
	"github.com/ryantrue/contractKeeper.git/internal/data"
	"net/http"
)

// LoginHandler обрабатывает запросы на вход в систему.
func LoginHandler(w http.ResponseWriter, r *http.Request, localizer *i18n.Localizer) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if data.ValidateUserCredentials(username, password) {
		token, err := auth.GenerateToken(username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/secure", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
	}
}
