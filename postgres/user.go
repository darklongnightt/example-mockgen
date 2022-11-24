package postgres

import (
	"database/sql"
	"example-mockgen/models"
	"fmt"
)

// UserRepo contains database dependencies
type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Insert creates an user entry
func (u *UserRepo) Insert(user *models.User) (*models.User, error) {
	fmt.Println("Insert user to db")

	return &models.User{
		Name: "User returned from db",
		Age:  123,
	}, nil
}

// Update updates an existing user
func (u *UserRepo) Update(user *models.User) (*models.User, error) {
	fmt.Println("Update user to db")

	return &models.User{
		Name: "User returned from db",
		Age:  123,
	}, nil
}
