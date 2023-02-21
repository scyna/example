package get_account

import (
	"ex/vertical_slide/shared/proto"
	"ex/vertical_slide/shared/repository"

	scyna "github.com/scyna/core"
)

const Path = "/account/get"

func Handler(ctx *scyna.Context, request *proto.GetAccountByEmailRequest) scyna.Error {
	ctx.Info("Receive GetAccountRequest")

	repository := repository.NewBaseRepository(ctx)
	account, err := repository.GetAccount(request.Email)
	if err != nil {
		return err
	}

	return ctx.OK(&proto.GetAccountResponse{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	})
}
