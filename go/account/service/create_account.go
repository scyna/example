package service

import (
	"ex/account/domain"
	"ex/account/model"

	scyna "github.com/scyna/core"

	proto "ex/account/proto"
)

func CreateAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Info("Receive CreateAccountRequest")

	service := domain.NewAccountService(ctx)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	if ret = service.AssureAccountNotExists(email); ret != nil {
		return ret
	}

	account := model.Account{
		LOG:   ctx,
		ID:    scyna.ID.Next(),
		Email: email,
		Name:  request.Name, /*TODO: check name*/
	}

	if account.Password, ret = model.ParsePassword(request.Password); ret != nil {
		ctx.Error("wrong password")
		return ret
	}

	command := scyna.NewCommand(ctx).
		SetEntity(account.ID).
		SetChannel(ACCOUNT_CREATED_CHANNEL).
		SetEvent(&proto.AccountCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email.String()})

	if ret = service.Repository.CreateAccount(command, &account); ret != nil {
		return ret
	}

	if ret = command.Commit(); ret != nil {
		return ret
	}

	return ctx.OK(&proto.CreateAccountResponse{Id: account.ID})
}
