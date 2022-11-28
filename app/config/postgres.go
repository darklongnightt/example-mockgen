package config

import (
	"fmt"
	"time"
)

// Postgres config struct
type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DBName   string `envconfig:"POSTGRES_DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"POSTGRES_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"POSTGRES_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"POSTGRES_MAX_IDLE_CONNECTION" default:"10"`
}

// ConnectionString builder based on postgres env values
func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.DBName)
}
