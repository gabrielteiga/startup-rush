package routes

import (
	"net/http"
)

type RouteInterface interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
