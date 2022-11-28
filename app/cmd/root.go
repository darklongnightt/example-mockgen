package cmd

import (
	"example-mockgen/app/config"
	"example-mockgen/models"
	"example-mockgen/postgres"
	"example-mockgen/s3"
	"example-mockgen/user"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import postgres driver
)

var (
	rootConfig *config.Root
	db         *sqlx.DB
)

func initConfig() {
	rootConfig = config.Init()
}

func initDB(pg *config.Postgres) {
	dbConn, err := sqlx.Open("postgres", pg.ConnectionString())
	if err != nil {
		log.Fatal("connection err: ", err)
	}

	dbConn.SetConnMaxLifetime(pg.MaxConnectionLifetime)
	dbConn.SetMaxOpenConns(pg.MaxOpenConnection)
	dbConn.SetMaxIdleConns(pg.MaxIdleConnection)

	if err = dbConn.Ping(); err != nil {
		log.Fatal("ping err: ", err)
	}
	db = dbConn
}

// Execute will call the root command execute
func Execute() {
	initConfig()
	initDB(rootConfig.Postgres)

	// Init the dependencies
	repo := postgres.New(db)
	s3 := s3.New()

	// Inject the dependencies into user service
	user := user.New(repo, s3)

	// Use functions
	user.AddUser(&models.User{
		Name: "user from main",
		Age:  12,
	})
}
