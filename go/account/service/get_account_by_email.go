package service

import (
	"ex/account/domain"
	"ex/account/model"

	proto "ex/account/proto/generated"

	scyna "github.com/scyna/core"
)

func GetAccountByEmailHandler(ctx *scyna.Endpoint, request *proto.GetAccountByEmailRequest) scyna.Error {
	ctx.Logger.Info("Receive GetUserRequest")

	repository := domain.LoadRepository(ctx.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	account, ret := repository.GetAccountByEmail(email)
	if ret != nil {
		return ret
	}

	ctx.Response(&proto.Account{
		Id:    account.ID,
		Email: account.Email.String(),
		Name:  account.Name,
		/*TODO*/
	})

	return scyna.OK
}
