package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reyesml/RMT/app/config"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/interactors"
	"github.com/reyesml/RMT/app/core/models"
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

	//Setup interactors
	createUser := interactors.NewCreateUser(userRepo)
	_ = createUser
	createSession := interactors.NewCreateSession(userRepo, sessionRepo, cfg.Session.SigningSecret)

	//Setup controllers
	authController := controllers.NewAuthController(createSession)

	//Setup routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(rmtMiddleware.Authenticate(sessionRepo, cfg.Session.SigningSecret))

		//This is just a sample route to demonstrate retrieving session info.
		//TODO: delete this after we have a useful authenticated route.
		r.Get("/session.info", func(w http.ResponseWriter, r *http.Request) {
			session, ok := r.Context().Value(models.SessionContextKey).(models.SessionContext)
			if !ok {
				w.Write([]byte("Couldn't locate session info :("))
				return
			}
			resp, err := json.Marshal(session)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(resp)
		})
	})

	//Unauthenticated routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world."))
	})
	r.Post("/login", authController.Login)

	http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), r)
}
