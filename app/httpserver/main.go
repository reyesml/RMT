package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reyesml/RMT/app/config"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/interactors"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/httpserver/controllers"
	rmtMiddleware "github.com/reyesml/RMT/app/httpserver/middleware"
	"net/http"
	"os"
)

func main() {
	configFile := os.Args[1]
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}
	//Setup repositories
	db, err := database.Connect(cfg.Database.DbId)
	if err != nil {
		panic(err)
	}
	userRepo := repos.NewUserRepo(db)
	sessionRepo := repos.NewSessionRepo(db)
	journalRepo := repos.NewJournalRepo(db)

	//Setup interactors
	createSession := interactors.NewCreateSession(userRepo, sessionRepo, cfg.Session.SigningSecret)

	//Setup controllers
	authController := controllers.NewAuthController(createSession)

	//Setup routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Unauthenticated routes
	r.Group(func(r chi.Router) {
		r.Post("/login", authController.Login)
	})

	//Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(rmtMiddleware.Authenticate(sessionRepo, cfg.Session.SigningSecret))

		cje := interactors.NewCreateJournal(journalRepo)
		journalController := controllers.NewJournalController(cje)
		r.Post("/journal", journalController.Create)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), r)
}
