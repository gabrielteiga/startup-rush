package routes

import (
	"github.com/gabrielteiga/startup-rush/api/controllers/health"
	"github.com/gabrielteiga/startup-rush/api/controllers/startup_controller"
	"github.com/gabrielteiga/startup-rush/api/controllers/tournament_controller"
	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/tournament_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/services"
	"github.com/gabrielteiga/startup-rush/internal/infra/repositories/gorm_adapter"
)

var dbConn = database.InitConnection()

var StartupRepository startup_entity.IStartupRepository = gorm_adapter.NewStartupGORMRepository(dbConn.DB)
var StartupService *services.StartupService = services.NewStartupService(StartupRepository)

var BattleRepository battle_entity.IBattleRepository = gorm_adapter.NewBattleGORMRepository(dbConn.DB)
var ParticipationRepository participations_entity.IParticipationRepository = gorm_adapter.NewParticipationsGORMRepository(dbConn.DB)

var TournamentRepository tournament_entity.ITournamentRepository = gorm_adapter.NewTournamentGORMRepository(dbConn.DB, ParticipationRepository, BattleRepository)
var TournamentService *services.TournamentService = services.NewTournamentService(TournamentRepository, StartupRepository, BattleRepository, ParticipationRepository)

func Provide() map[string]RouteInterface {
	dbConn.Migrate()

	return map[string]RouteInterface{
		"GET /api/health": health.NewHealthController(),

		"POST /api/startups": startup_controller.NewCreateStartup(StartupService),
		"GET /api/startups":  startup_controller.NewListStartup(StartupService),

		"POST /api/tournaments":          tournament_controller.NewCreateTournament(TournamentService),
		"GET /api/tournaments":           tournament_controller.NewListTournament(TournamentService),
		"GET /api/tournaments/":          tournament_controller.NewGetTournamentByID(TournamentService),
		"GET /api/tournaments/start/":    tournament_controller.NewStartTournamentByID(TournamentService),
		"GET /api/tournaments/startups/": startup_controller.NewGetStartupsByTournamentID(TournamentService, StartupService),
	}
}
