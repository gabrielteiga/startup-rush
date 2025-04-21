package main

import (
	"log"

	"github.com/gabrielteiga/startup-rush/database"
)

func main() {
	conn := database.InitConnection()
	if conn.Error != nil {
		log.Fatalln(conn.Error)
	}

	err := conn.Migrate()
	if err != nil {
		log.Fatalln(err)
	}
}
