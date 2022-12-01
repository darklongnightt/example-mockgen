package postgres

import (
	"example-mockgen/models"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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

// AddUser creates an user entry
func (r *Repository) AddUser(user *models.User) (*models.User, error) {
	query := "INSERT INTO users (id, name, age, created_at) VALUES ($1, $2, $3, $4)"
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	user.ID = uuid.String()
	user.CreatedAt = time.Now()

	res, err := r.db.Exec(query, user.ID, user.Name, user.Age, user.CreatedAt)
	if err != nil {
		return nil, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers get all users
func (r *Repository) GetUsers() ([]*models.User, error) {
	builder := squirrel.Select("id", "name", "age", "created_at").PlaceholderFormat(squirrel.Dollar)
	builder = builder.OrderBy("created_at DESC").From("users")

	sqlStr, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []*models.User{}
	for rows.Next() {
		u := new(models.User)
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}
