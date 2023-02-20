package create_account

import (
	"ex/account/proto"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
)

const URL = "/account/create"

var (
	ACCOUNT_EXISTED     = scyna.NewError(100, "User Existed")
	ACCOUNT_NOT_EXISTED = scyna.NewError(101, "User NOT Existed")
)

func Handler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive CreateAccountRequest")

	if validateCreateAccountRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	repository := LoadRepository(&ctx.Context)

	if _, err := repository.GetAccount(request.Email); err == nil {
		return ACCOUNT_EXISTED
	}

	account := &Account{
		ID:       scyna.ID.Next(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := repository.CreateAccount(account); err != nil {
		return err
	}

	return ctx.OK(&proto.CreateAccountResponse{Id: account.ID})
}

func validateCreateAccountRequest(request *proto.CreateAccountRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&request.Password, validation.Required, validation.Length(1, 18)),
	)
}
