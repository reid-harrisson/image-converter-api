package requests

import validation "github.com/go-ozzo/ozzo-validation"

type ImageRequest struct {
	BackColor string `json:"back_color" example:"#000000"`
	ForeColor string `json:"fore_color" example:"#FFFFFF"`
}

func (request ImageRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.BackColor, validation.Required),
		validation.Field(&request.ForeColor, validation.Required),
	)
}
