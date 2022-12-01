package postgres

import (
	"example-mockgen/app/config"

	"github.com/jmoiron/sqlx"
)

// DB returns instance of db connection
func DB(pg *config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", pg.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(pg.MaxConnectionLifetime)
	db.SetMaxOpenConns(pg.MaxOpenConnection)
	db.SetMaxIdleConns(pg.MaxIdleConnection)

	return db, db.Ping()
}
