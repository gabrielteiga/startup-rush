package configs

import (
	"os"
	"time"
)

var (
	// APP VARIABLES
	APP_PORT = os.Getenv("APP_PORT")

	// DATABASE VARIABLES
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DOMAIN   = os.Getenv("DB_DOMAIN")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_DATABASE = os.Getenv("DB_DATABASE")

	// TIME LAYOUT ISO8601 ("2021-10-20T15:04:05Z")
	DATE_LAYOUT = time.RFC3339
)
