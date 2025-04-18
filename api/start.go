package api

import (
	"log"
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/routes"
)

const SERVER_PORT = ":8080"

func Start() {
	mux := http.NewServeMux()

	initEndpoints(mux, routes.Provide())

	log.Printf("Starting server in the %s port...", SERVER_PORT)
	log.Fatal(http.ListenAndServe(SERVER_PORT, mux))
}

func initEndpoints(mux *http.ServeMux, controllers map[string]routes.RouteInterface) {
	for endpoint, controller := range controllers {
		mux.HandleFunc(endpoint, controller.Handle)
	}
}
