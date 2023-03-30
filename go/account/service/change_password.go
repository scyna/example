package service

import (
	"ex/account/domain"
	proto "ex/account/proto"

	scyna "github.com/scyna/core"
)

func ChangePasswordHandler(ctx *scyna.Endpoint, request *proto.ChangePasswordRequest) scyna.Error {
	ctx.Info("Receive ChangePasswordRequest")

	service := domain.NewAccountService(ctx)
	account, ret := service.Repository.GetAccountByID(request.Id)
	if ret != nil {
		return ret
	}

	if ret = service.Repository.LoadPassword(account); ret != nil {
		return ret
	}

	if ret = account.ChangePassword(request.Current, request.Future); ret != nil {
		return ret
	}

	command := scyna.NewCommand(ctx).
		SetEntity(account.ID).
		SetChannel(PASSWORD_CHANGED_CHANNEL).
		SetEvent(&proto.PasswordChanged{
			Id:      account.ID,
			Current: request.Current,
			Future:  request.Future})

	if ret = service.Repository.UpdatePassword(command, account); ret != nil {
		return ret
	}

	if ret = command.Commit(); ret != nil {
		return ret
	}

	return scyna.OK
}
