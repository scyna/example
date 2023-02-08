package main

import (
	"ex/account/domain"
	"ex/account/repository"
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

	domain.AttachRepositoryCreator(repository.NewRepository)
	scyna.RegisterEndpoint(service.GET_ACCOUNT_URL, service.GetAccountByEmailHandler)
	scyna.Start()
}
