package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/interactors"
	"github.com/reyesml/RMT/app/repos"
	"github.com/reyesml/RMT/app/server/controllers"
	"net/http"
)

func main() {
	//Setup repositories
	db, err := database.Connect("dev.db")
	if err != nil {
		panic(err)
	}
	userRepo := repos.NewUserRepo(db)
	sessionRepo := repos.NewSessionRepo(db)

	//Setup interactors
	createUser := interactors.CreateUser{UserRepo: userRepo}
	_ = createUser
	createSession := interactors.CreateSession{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}

	//Setup controllers
	authController := controllers.NewAuthController(createSession)

	//Setup routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world."))
	})

	r.Post("/login", authController.Login)

	http.ListenAndServe(":3000", r)
}
