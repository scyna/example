package service

import (
	"ex/transaction_script/proto"

	scyna "github.com/scyna/core"
)

func GetAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive GetAccountRequest")
	return nil
}
