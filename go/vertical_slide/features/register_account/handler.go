package register_account

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"

	scyna "github.com/scyna/core"
)

const URL = "/account/create"

func Handler(ctx *scyna.Endpoint, request *proto.RegisterAccountRequest) scyna.Error {
	ctx.Info("Receive CreateAccountRequest")

	if validateRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	repository := NewRepository(&ctx.Logger)

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

	return ctx.OK(&proto.RegisterAccountResponse{ID: account.ID})
}
