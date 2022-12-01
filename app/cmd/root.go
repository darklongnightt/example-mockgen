package cmd

import (
	"example-mockgen/app/config"
	"example-mockgen/models"
	"example-mockgen/postgres"
	"example-mockgen/s3"
	"example-mockgen/user"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import postgres driver
)

// Execute will call the root command execute
func Execute() {
	// Init all dependencies
	cfg := config.New()
	db, err := postgres.DB(cfg.Postgres)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	repo := postgres.New(db)
	s3 := s3.New()

	// Inject the dependencies into user service
	user := user.New(repo, s3)

	// Use functions
	if _, err = user.AddUser(&models.User{
		Name: "user from main",
		Age:  12,
	}); err != nil {
		log.Fatal("AddUser error: ", err)
	}

	users, err := user.GetUsers()
	if err != nil {
		log.Fatal("GetUsers error: ", err)
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
