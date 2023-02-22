package event

import (
	"log"

	proto "ex/account/proto"

	scyna "github.com/scyna/core"
)

func SendEmailHandler(ctx scyna.Context, event *proto.SendEmailTask) {
	ctx.Info("Receive SendEmailTask")
	log.Print(event)
}
