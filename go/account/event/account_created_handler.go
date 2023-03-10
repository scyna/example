package event

import (
	"log"
	"time"

	proto "ex/account/proto"
	"ex/account/service"

	scyna "github.com/scyna/core"
)

func AccountCreatedHandler(ctx *scyna.Event, event *proto.AccountCreated) {
	ctx.Info("AccountCreated handler")
	log.Print(event)
	ctx.ScheduleTask(service.SEND_EMAIL_CHANNEL, time.Now(), 61, &proto.SendEmailTask{
		Email:   event.Email,
		Content: "Just say hello",
	}, 2)
}
