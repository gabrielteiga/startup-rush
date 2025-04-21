package startup_entity

import "time"

type Startup struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name"`
	Slogan     string    `json:"slogan"`
	Foundation time.Time `json:"foundation"`
}

func NewStartup(id uint, name, slogan string, foundation time.Time) *Startup {
	return &Startup{
		ID:         id,
		Name:       name,
		Slogan:     slogan,
		Foundation: foundation,
	}
}
