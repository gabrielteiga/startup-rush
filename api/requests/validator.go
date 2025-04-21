package requests

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Validate(request interface{}) *ErrorResponse {
	err := validate.Struct(request)
	if err != nil {
		return &ErrorResponse{
			Error: err.Error(),
		}
	}
	return nil
}
