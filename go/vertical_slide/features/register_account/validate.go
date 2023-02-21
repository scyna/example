package register_account

import (
	"ex/vertical_slide/shared/proto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func validateRequest(request *proto.RegisterAccountRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&request.Password, validation.Required, validation.Length(1, 18)),
	)
}
