package config

// App config struct
type App struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"sample-service"`
	Env         string `envconfig:"APP_ENV" default:"development"`
}
