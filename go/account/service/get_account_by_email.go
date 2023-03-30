package service

import (
	"ex/account/domain"
	"ex/account/model"

	proto "ex/account/proto"

	scyna "github.com/scyna/core"
)

func GetAccountByEmailHandler(ctx *scyna.Endpoint, request *proto.GetAccountByEmailRequest) scyna.Error {
	ctx.Info("Receive GetAccountRequest")

	service := domain.NewAccountService(ctx)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	account, ret := service.Repository.GetAccountByEmail(email)
	if ret != nil {
		return ret
	}

	return ctx.OK(&proto.GetAccountResponse{
		Id:    account.ID,
		Email: account.Email.String(),
		Name:  account.Name,
		/*TODO*/
	})
}
