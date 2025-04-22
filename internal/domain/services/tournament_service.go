package services

import (
	"log"
	"math/rand/v2"

	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/battle_events_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/event_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/participations_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/startup_entity"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/tournament_entity"
)

type TournamentService struct {
	TournamentRepository     tournament_entity.ITournamentRepository
	StartupRepository        startup_entity.IStartupRepository
	BattleRepository         battle_entity.IBattleRepository
	ParticipationsRepository participations_entity.IParticipationRepository
	EventRepository          event_entity.IEventRepository
	BattleEventRepository    battle_events_entity.IBattleEventsRepository
}

func NewTournamentService(
	tournamentRepository tournament_entity.ITournamentRepository,
	startupRepository startup_entity.IStartupRepository,
	battleRepository battle_entity.IBattleRepository,
	participationsRepository participations_entity.IParticipationRepository,
	eventRepository event_entity.IEventRepository,
	battleEventRepository battle_events_entity.IBattleEventsRepository,
) *TournamentService {
	return &TournamentService{
		TournamentRepository:     tournamentRepository,
		StartupRepository:        startupRepository,
		BattleRepository:         battleRepository,
		ParticipationsRepository: participationsRepository,
		EventRepository:          eventRepository,
		BattleEventRepository:    battleEventRepository,
	}
}

func (ts *TournamentService) Create(startupIDs []uint) *tournament_entity.Tournament {
	startups := ts.StartupRepository.FindByIDs(startupIDs)
	if len(startups) != len(startupIDs) {
		return nil
	}

	tournament, err := ts.TournamentRepository.Create(startups)
	if err != nil {
		log.Println("Error creating tournament:", err)
		return nil
	}
	return tournament
}

func (ts *TournamentService) List() []*tournament_entity.Tournament {
	tournaments, err := ts.TournamentRepository.List()
	if err != nil {
		log.Println("Error listing tournaments:", err)
		return nil
	}

	return tournaments
}

func (ts *TournamentService) Start(tournamentID uint) *tournament_entity.Tournament {
	tournament, err := ts.TournamentRepository.FindByID(tournamentID)
	if err != nil {
		log.Println("Error finding tournament: ", err)
		return nil
	}

	if tournament.Finished || len(tournament.Battles) > 0 {
		log.Println("Tournament already finished or battles already generated")
		return tournament_entity.NewTournament(tournament.ID, tournament.Finished, tournament.ChampionID, tournament.Participants, tournament.Battles)
	}

	if err := ts.GenerateBattles(tournament); err != nil {
		log.Println("Error generating battles: ", err)
		return nil
	}

	return tournament_entity.NewTournament(tournament.ID, tournament.Finished, tournament.ChampionID, tournament.Participants, tournament.Battles)
}

func (ts *TournamentService) GenerateBattles(tournament *tournament_entity.Tournament) error {
	participants := tournament.Participants

	rand.Shuffle(len(participants), func(i, j int) {
		participants[i], participants[j] = participants[j], participants[i]
	})

	var phase battle_entity.BattlePhase

	switch len(participants) {
	case 4:
		phase = battle_entity.PhaseSemiFinal
	case 8:
		phase = battle_entity.PhaseQuarterFinal
	}

	var battles []*battle_entity.Battle
	for i := 0; i < len(participants); i = i + 2 {
		p1 := participants[i]
		p2 := participants[i+1]

		battle, err := ts.BattleRepository.Create(tournament.ID, p1.StartupID, p2.StartupID, nil, nil, phase)
		if err != nil {
			return err
		}

		battles = append(battles, battle)
	}
	tournament.Battles = battles
	return nil
}

func (ts *TournamentService) GetByID(id uint) *tournament_entity.Tournament {
	tournament, err := ts.TournamentRepository.FindByID(id)
	if err != nil {
		log.Println("Error finding tournament:", err)
		return nil
	}

	return tournament
}

func (ts *TournamentService) FindParticipantsByTournamentID(tournamentID uint) []*participations_entity.Participation {
	startups, err := ts.ParticipationsRepository.FindByTournamentID(tournamentID)
	if err != nil {
		log.Println("Error finding participants:", err)
		return nil
	}

	return startups
}

func (ts *TournamentService) GetBattleByID(id uint) (*battle_entity.Battle, error) {
	battle, err := ts.BattleRepository.FindByID(id)
	if err != nil {
		log.Println("Error finding battle:", err)
		return nil, err
	}
	return battle, nil
}

func (ts *TournamentService) GetEvents() ([]*event_entity.Event, error) {
	events, err := ts.EventRepository.List()
	if err != nil {
		log.Println("Error finding events:", err)
		return nil, err
	}
	return events, nil
}

func (ts *TournamentService) Battle(battleID uint, events map[uint][]uint) (any, error) {
	battle, err := ts.BattleRepository.FindByID(battleID)
	if err != nil {
		log.Println("Error finding battle:", err)
		return nil, err
	}

	err = ts.registerAndCheckEvents(battle, events)
	if err != nil {
		log.Println("Error registering events:", err)
		return nil, err
	}

	// We are using here the gorm model. Need to change this to the entity model.
	battleDatabase, err := ts.BattleEventRepository.GetBattleDatabaseWithEvents(battle.ID)
	if err != nil {
		log.Println("Error finding battle:", err)
		return nil, err
	}

	return battleDatabase, nil

}

func (ts *TournamentService) registerAndCheckEvents(battle *battle_entity.Battle, events map[uint][]uint) error {
	for startupID, eventsIDs := range events {
		for _, eventID := range eventsIDs {
			_, err := ts.BattleEventRepository.Create(battle.ID, startupID, eventID)
			if err != nil {
				log.Println("Error creating battle event:", err)
				return err
			}
		}
	}
	return nil
}
