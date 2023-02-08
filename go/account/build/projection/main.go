package main

import (
	"ex/account/event"
	"ex/account/service"

	scyna "github.com/scyna/core"
)

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       "ex_account",
		Secret:     "123456",
	})
	defer scyna.Release()

	scyna.RegisterEvent("ex_account", service.ACCOUNT_CREATED_CHANNEL, event.AccountCreatedHandler)
	scyna.RegisterTask("ex_account", service.SEND_EMAIL_CHANNEL, event.SendEmailHandler)
	scyna.Start()
}
