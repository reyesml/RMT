package main

import (
	"fmt"
	"github.com/reyesml/RMT/app/core/config"
	"github.com/reyesml/RMT/app/core/database"
	"github.com/reyesml/RMT/app/repos"
	"os"
)

func main() {
	configFile := os.Args[1]
	cfg, err := config.LoadConfig(configFile)
	db, err := database.Connect(cfg.Database.DbId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Performing db migrations...")
	err = repos.MigrateAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations complete.")

	userRepo := repos.NewUserRepo(db)
	admin, err := userRepo.GetByUsername("admin")
	if err == nil {
		fmt.Println("Admin user exists. Exiting.")
		return
	}

	fmt.Println("Creating Admin user...")
	admin.Username = "admin"
	admin.Admin = true
	err = admin.SetPassword("not_secure")
	if err != nil {
		panic(err)
	}
	err = userRepo.Create(&admin)
	if err != nil {
		panic(err)
	}
	fmt.Println("Admin user created.")
}
