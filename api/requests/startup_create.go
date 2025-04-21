package requests

type RequestStartupCreate struct {
	Name       string `json:"name" validate:"required"`
	Slogan     string `json:"slogan" validate:"required"`
	Foundation string `json:"foundation" validate:"required"`
}

func NewRequestStartupCreate(name, slogan, foundation string) *RequestStartupCreate {
	return &RequestStartupCreate{
		Name:       name,
		Slogan:     slogan,
		Foundation: foundation,
	}
}
