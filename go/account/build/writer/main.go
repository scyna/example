package main

import (
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

	scyna.InitSingleWriter("ex_account")
	scyna.RegisterEndpoint(service.CREATE_ACCOUNT_URL, service.CreateAccountHandler)
	scyna.Start()
}
