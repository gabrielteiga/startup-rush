package requests

type RequestCreateTournament struct {
	StartupsIDs []uint `json:"startupsIds" validate:"required,eq=4|eq=8"`
}

func NewRequestCreateTournament(startupIDs []uint) *RequestCreateTournament {
	return &RequestCreateTournament{
		StartupsIDs: startupIDs,
	}
}
