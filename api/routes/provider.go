package routes

import (
	"github.com/gabrielteiga/startup-rush/api/controllers/health"
)

func Provide() map[string]RouteInterface {
	return map[string]RouteInterface{
		"/api/health": health.NewHealthController(),
	}
}
