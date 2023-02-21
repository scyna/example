package register_account

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"

	scyna "github.com/scyna/core"
)

const Path = "/account/register"

func Handler(ctx *scyna.Context, request *proto.RegisterAccountRequest) scyna.Error {
	ctx.Info("Receive RegisterAccountRequest")

	if validateRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	repository := NewRepository(ctx)

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
