package postgres

import (
	"example-mockgen/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Repository contains database dependencies and implements methods to interact with db
type Repository struct {
	db *sqlx.DB
}

// New will return
func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// InsertUser creates an user entry
func (u *Repository) InsertUser(user *models.User) (*models.User, error) {
	fmt.Println("InsertUser user to db")

	return &models.User{
		Name: "User returned from db",
		Age:  123,
	}, nil
}

// UpdateUser updates an existing user
func (u *Repository) UpdateUser(user *models.User) (*models.User, error) {
	fmt.Println("Update user to db")

	return &models.User{
		Name: "User returned from db",
		Age:  123,
	}, nil
}
