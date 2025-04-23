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

	handler := corsMiddleware(mux)

	log.Printf("Starting server in the %s port...", PORT)
	log.Fatal(http.ListenAndServe(PORT, handler))
}

func initEndpoints(mux *http.ServeMux, controllers map[string]routes.RouteInterface) {
	for endpoint, controller := range controllers {
		mux.HandleFunc(endpoint, controller.Handle)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
