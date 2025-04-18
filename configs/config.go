package configs

import "os"

var (
	PORT = os.Getenv("PORT")
)
