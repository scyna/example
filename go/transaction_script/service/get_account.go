package service

import (
	"ex/transaction_script/proto"
	"ex/transaction_script/repository"

	scyna "github.com/scyna/core"
)

func GetAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive GetAccountRequest")

	repository := repository.LoadRepository(&ctx.Context)
	account, err := repository.GetAccount(request.Email)
	if err != nil {
		return err
	}

	ctx.Response(&proto.GetAccountResponse{
		Id:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	})
	return nil
}
