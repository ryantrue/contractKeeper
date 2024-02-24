package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ryantrue/contractKeeper.git/internal/config"
	"github.com/ryantrue/contractKeeper.git/internal/handlers"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"path/filepath"
)

func initLocalization() *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(filepath.Join("locales", "ru.json"))
	localizer := i18n.NewLocalizer(bundle, "ru")
	welcomeMessage, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: "Welcome"})
	if err != nil {
		log.Fatalf("Failed to localize: %v", err)
	}
	log.Println(welcomeMessage) // Логгирование приветственного сообщения
	return localizer
}
func setupRouter(localizer *i18n.Localizer) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomePageHandler(w, r, localizer)
	}).Methods("GET")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, localizer)
	}).Methods("POST")
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, localizer)
	}).Methods("GET", "POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return router
}
func main() {
	config.LoadConfig()
	localizer := initLocalization()
	router := setupRouter(localizer)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + config.AppConfig.Port,
		// дополнительные параметры сервера, такие как WriteTimeout и ReadTimeout
	}
	log.Printf("Starting server on port %s", config.AppConfig.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
