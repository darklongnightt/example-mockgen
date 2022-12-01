package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Root config
type Root struct {
	App      *App
	Postgres *Postgres
}

// New inits the root configuration by loading variables
// from the environment, plus the filenames provided
func New(filenames ...string) *Root {
	// Ignore if there is no .env file.
	_ = godotenv.Overload(filenames...)

	r := new(Root)
	if err := envconfig.Process("", r); err != nil {
		log.Fatal(err)
	}
	return r
}
