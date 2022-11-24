package main

import (
	"database/sql"
	"example-mockgen/models"
	"example-mockgen/postgres"
	"example-mockgen/s3"
	"example-mockgen/services"
)

func main() {
	// Init the dependencies
	userRepo := postgres.NewUserRepo(&sql.DB{})
	s3 := s3.NewClient()

	// Inject the dependencies into user service
	userService := services.NewUserService(userRepo, s3)

	// Use functions
	userService.CreateUser(&models.User{})
}
