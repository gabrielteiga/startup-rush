package api

import (
	"log"
	"net/http"

	"github.com/gabrielteiga/startup-rush/api/routes"
	"github.com/gabrielteiga/startup-rush/configs"
)

var PORT = ":" + configs.APP_PORT

func Start() {
	mux := http.NewServeMux()

	initEndpoints(mux, routes.Provide())

	log.Printf("Starting server in the %s port...", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}

func initEndpoints(mux *http.ServeMux, controllers map[string]routes.RouteInterface) {
	for endpoint, controller := range controllers {
		mux.HandleFunc(endpoint, controller.Handle)
	}
}
