package main

import (
	"database/sql"
	"example-mockgen/models"
	"example-mockgen/postgres"
	"example-mockgen/s3"
	"example-mockgen/user"
)

func main() {
	// Init the dependencies
	repo := postgres.New(&sql.DB{})
	s3 := s3.NewClient()

	// Inject the dependencies into user service
	user := user.New(repo, s3)

	// Use functions
	user.CreateUser(&models.User{})
}
