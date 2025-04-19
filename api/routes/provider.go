package routes

import (
	"github.com/gabrielteiga/startup-rush/api/controllers/health"
)

func Provide() map[string]RouteInterface {
	return map[string]RouteInterface{
		"GET /api/health": health.NewHealthController(),
	}
}
