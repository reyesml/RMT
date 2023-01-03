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
	personRepo := repos.NewPersonRepo(db)
	personQualityRepo := repos.NewPersonQualityRepo(db)
	qualityRepo := repos.NewQualityRepo(db)
	noteRepo := repos.NewNoteRepo(db)

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

		r.Route("/journal", func(r chi.Router) {
			journalController := controllers.NewJournalController(
				interactors.NewCreateJournal(journalRepo),
				interactors.NewGetJournal(journalRepo),
				interactors.NewListJournals(journalRepo),
			)
			r.Get("/", journalController.List)
			r.Post("/", journalController.Create)
			r.Get("/{UUID}", journalController.Get)
		})

		r.Route("/people", func(r chi.Router) {
			personController := controllers.NewPersonController(
				interactors.NewCreatePerson(personRepo),
				interactors.NewGetPerson(personRepo),
				interactors.NewCreatePersonQuality(personRepo, qualityRepo, personQualityRepo),
				interactors.NewListPersonQualities(personRepo, personQualityRepo),
				interactors.NewCreatePersonNote(personRepo, noteRepo),
				interactors.NewListPersonNotes(personRepo, noteRepo),
				interactors.NewListPeople(personRepo),
			)
			r.Get("/", personController.List)
			r.Post("/", personController.Create)
			r.Get("/{UUID}", personController.Get)
			r.Get("/{UUID}/qualities", personController.ListPersonQualities)
			r.Post("/{UUID}/qualities", personController.CreatePersonQuality)
			r.Get("/{UUID}/notes", personController.ListNotes)
			r.Post("/{UUID}/notes", personController.CreateNote)
		})

		r.Route("/person-quality", func(r chi.Router) {
			personQualityController := controllers.NewPersonQualityController(
				interactors.NewGetPersonQuality(personQualityRepo),
				interactors.NewCreatePersonQualityNote(personQualityRepo, noteRepo),
				interactors.NewListPersonQualityNotes(personQualityRepo, noteRepo),
			)
			r.Get("/{UUID}", personQualityController.Get)
			r.Get("/{UUID}/notes", personQualityController.ListNotes)
			r.Post("/{UUID}/notes", personQualityController.CreateNote)
			//Delete
			//Edit
		})

		r.Route("/qualities", func(r chi.Router) {
			qualityController := controllers.NewQualityController(
				interactors.NewListQualities(qualityRepo),
			)
			r.Get("/", qualityController.List)
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), r)
}
