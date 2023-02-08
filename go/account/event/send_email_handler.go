package event

import (
	"log"

	proto "ex/account/proto/generated"

	scyna "github.com/scyna/core"
)

func SendEmailHandler(ctx *scyna.Context, event *proto.SendEmailTask) {
	ctx.Logger.Info("Receive SendEmailTask")
	log.Print(event)
}
