package main

import (
	"fmt"
	"github.com/reyesml/RMT/app/config"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/core/models"
	"github.com/reyesml/RMT/app/core/repos"
	"github.com/reyesml/RMT/app/core/utils"
	"os"
)

func main() {
	configFile := os.Args[1]
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}
	db, err := database.Connect(cfg.Database.DbId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Performing db migrations...")
	err = utils.MigrateAllModels(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations complete.")

	userRepo := repos.NewUserRepo(db)
	_, err = userRepo.GetByUsername("admin")
	if err == nil {
		fmt.Println("Admin user exists. Exiting.")
		return
	}

	fmt.Println("Creating Admin user...")
	admin, err := models.NewUser("admin", "not_secure")
	admin.Admin = true
	if err != nil {
		panic(err)
	}
	err = userRepo.Create(admin)
	if err != nil {
		panic(err)
	}
	fmt.Println("Admin user created.")
}
