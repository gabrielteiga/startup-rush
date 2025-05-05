package services

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/gabrielteiga/startup-rush/database"
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

	if battle.Finished {
		log.Println("Battle already finished")
		return nil, fmt.Errorf("battle already finished")
	}

	err = ts.registerAndCheckEvents(battle, events)
	if err != nil {
		log.Println("Error registering events:", err)
		return nil, err
	}

	// battleDatabase -> We are using here the gorm model. Need to change this to the entity model.
	battleDatabase, err := ts.BattleEventRepository.GetBattleDatabaseWithEvents(battle.ID)
	if err != nil {
		log.Println("Error finding battle:", err)
		return nil, err
	}

	err = ts.defineBattleWinner(battleDatabase)
	if err != nil {
		log.Println("Error defining battle winner:", err)
		return nil, err
	}

	if err := ts.ParticipationsRepository.AddScore(battleDatabase.TournamentID, *battleDatabase.WinnerID, 30); err != nil {
		return nil, fmt.Errorf("error adding victory bonus: %w", err)
	}

	if err := ts.advancePhaseIfNeeded(battleDatabase); err != nil {
		return nil, fmt.Errorf("error advancing phase: %w", err)
	}

	BattleUpdated, err := ts.BattleRepository.FindByID(battleID)
	if err != nil {
		return nil, fmt.Errorf("error reloading battle: %w", err)
	}

	return BattleUpdated, nil
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

// Using GORM Model
func (ts *TournamentService) defineBattleWinner(battleDatabase *database.Battle) error {
	score1, score2 := ts.calculateScores(battleDatabase)

	battleDatabase.ScoreStartup1 = &score1
	battleDatabase.ScoreStartup2 = &score2
	battleDatabase.Finished = true

	var winnerID uint
	if score1 > score2 {
		winnerID = battleDatabase.Startup1ID
	} else {
		winnerID = battleDatabase.Startup2ID
	}
	battleDatabase.WinnerID = &winnerID

	battleEntity := battle_entity.NewBattle(
		battleDatabase.ID,
		battleDatabase.TournamentID,
		battleDatabase.Startup1ID,
		battleDatabase.Startup2ID,
		battleDatabase.ScoreStartup1,
		battleDatabase.ScoreStartup2,
		battleDatabase.Finished,
		battleDatabase.WinnerID,
		battleDatabase.BattleParentID,
		battleDatabase.BattleChildren1ID,
		battleDatabase.BattleChildren2ID,
		battleDatabase.Phase,
	)
	if err := ts.BattleRepository.SaveBattle(battleEntity); err != nil {
		return fmt.Errorf("error saving battle result: %w", err)
	}
	return nil
}

// Using GORM Model
func (ts *TournamentService) calculateScores(battleDatabase *database.Battle) (int, int) {
	var score1, score2 int
	for _, event := range battleDatabase.BattleEvents {
		if event.StartupID == battleDatabase.Startup1ID {
			score1 += event.Event.Score
		} else {
			score2 += event.Event.Score
		}
	}

	// Shark Fight
	if score1 == score2 {
		if rand.IntN(2) == 0 {
			score1 += 2
		} else {
			score2 += 2
		}
	}

	return score1, score2
}

// Using GORM Model
func (ts *TournamentService) advancePhaseIfNeeded(GORMBattle *database.Battle) error {
	phase := GORMBattle.Phase
	tournamentID := GORMBattle.TournamentID

	pending, err := ts.BattleRepository.CountByPhase(tournamentID, phase, false)
	if err != nil {
		return err
	}
	if pending > 0 {
		return nil
	}

	if phase == battle_entity.PhaseFinal {
		_, err := ts.TournamentRepository.Finish(tournamentID, GORMBattle.WinnerID)
		return err
	}

	var nextPhase battle_entity.BattlePhase
	switch phase {
	case battle_entity.PhaseQuarterFinal:
		nextPhase = battle_entity.PhaseSemiFinal
	case battle_entity.PhaseSemiFinal:
		nextPhase = battle_entity.PhaseFinal
	}

	winnersBattlesMap, err := ts.BattleRepository.FindWinnersAndBattlesByPhase(tournamentID, phase)
	if err != nil {
		return err
	}
	for i := 0; i < len(winnersBattlesMap); i += 2 {
		_, err := ts.BattleRepository.Create(
			tournamentID,
			winnersBattlesMap[i].WinnerID,
			winnersBattlesMap[i+1].WinnerID,
			&winnersBattlesMap[i].BattleID,
			&winnersBattlesMap[i+1].BattleID,
			nextPhase,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

type EventCount struct {
	EventName string `json:"eventName"`
	Count     int    `json:"count"`
}

type RankingEntry struct {
	StartupID   uint         `json:"startupId"`
	Name        string       `json:"name"`
	Slogan      string       `json:"slogan"`
	Score       int          `json:"score"`
	EventCounts []EventCount `json:"events"`
	Battles     []*battle_entity.Battle `json:"battles"`
}

func (ts *TournamentService) GetRanking(tournamentID uint) ([]*RankingEntry, error) {
	participants, err := ts.ParticipationsRepository.FindRankingByTournamentID(tournamentID)
	if err != nil {
		log.Println("Error finding participants:", err)
		return nil, err
	}

	eventStats, err := ts.BattleEventRepository.CountEventsByTournament(tournamentID)
	if err != nil {
		log.Println("Error counting events:", err)
		return nil, err
	}

	eventMap := make(map[uint][]EventCount)
	for _, eventStat := range eventStats {
		eventMap[eventStat.StartupID] = append(eventMap[eventStat.StartupID], EventCount{
			EventName: eventStat.EventName,
			Count:     eventStat.Total,
		})
	}

	battles, err := ts.BattleRepository.FindByTournamentID(tournamentID)
	if err != nil {
		log.Println("Error finding battles:", err)
		return nil, err
	}

	battleMap := make(map[uint][]*battle_entity.Battle)
	for _, battle := range battles {
		battleMap[battle.Startup1ID] = append(battleMap[battle.Startup1ID], battle)
		battleMap[battle.Startup2ID] = append(battleMap[battle.Startup2ID], battle)
	}

	var ranking []*RankingEntry
	for _, participant := range participants {
		startup := ts.StartupRepository.FindByID(participant.StartupID)
		if startup == nil {
			log.Println("Error finding startup:", err)
			return nil, err
		}

		ranking = append(ranking, &RankingEntry{
			StartupID:   participant.StartupID,
			Name:        startup.Name,
			Slogan:      startup.Slogan,
			Score:       participant.Score,
			EventCounts: eventMap[participant.StartupID],
			Battles:     battleMap[participant.StartupID],
		})
	}
	return ranking, nil
}
