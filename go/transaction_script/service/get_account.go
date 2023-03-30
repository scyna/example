package service

import (
	"ex/transaction_script/proto"
	"ex/transaction_script/repository"

	scyna "github.com/scyna/core"
)

func GetAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Info("Receive GetAccountRequest")

	repository := repository.LoadRepository(ctx)
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
