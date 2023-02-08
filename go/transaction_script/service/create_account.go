package service

import (
	"ex/transaction_script/proto"
	"ex/transaction_script/repository"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
)

func CreateAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive CreateAccountRequest")
	db := repository.LoadRepository(&ctx.Context)

	if validateCreateAccountRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	if _, err := db.GetAccount(request.Email); err == nil {
		return repository.ACCOUNT_EXISTED
	}

	account := &repository.Account{
		ID:       scyna.ID.Next(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.CreateAccount(account); err != nil {
		return err
	}

	ctx.Response(&proto.CreateAccountResponse{Id: account.ID})

	return scyna.OK
}

func validateCreateAccountRequest(request *proto.CreateAccountRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Name, validation.Required, validation.Length(1, 150)),
		validation.Field(&request.Password, validation.Required, validation.Length(1, 18)),
	)
}
