package service

import (
	"ex/transaction_script/model"
	"ex/transaction_script/proto"
	"ex/transaction_script/repository"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
)

func CreateAccountHandler(ctx *scyna.Context, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Info("Receive CreateAccountRequest")
	repository := repository.LoadRepository(ctx)

	if validateCreateAccountRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetAccount(request.Email); err == nil {
		return model.ACCOUNT_EXISTED
	}

	account := &model.Account{
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
