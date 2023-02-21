package get_account

import (
	"ex/vertical_slide/shared/proto"
	"ex/vertical_slide/shared/repository"

	scyna "github.com/scyna/core"
)

const URL = "/account/get"

func Handler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive GetAccountRequest")

	repository := repository.NewBaseRepository(&ctx.Logger)
	account, err := repository.GetAccount(request.Email)
	if err != nil {
		return err
	}

	return ctx.OK(&proto.GetAccountResponse{
		Id:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	})
}
