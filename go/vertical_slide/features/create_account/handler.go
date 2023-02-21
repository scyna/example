package create_account

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"
	"ex/vertical_slide/shared/repository"

	scyna "github.com/scyna/core"
)

const URL = "/account/create"

func Handler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Info("Receive CreateAccountRequest")

	if validateRequest(request) != nil {
		return scyna.REQUEST_INVALID
	}

	repository := repository.LoadRepository(&ctx.Context)

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
