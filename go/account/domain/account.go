package domain

import (
	"ex/account/repository"

	scyna "github.com/scyna/core"
)

type accountService struct {
	Repository *repository.AccountRepository
	context    *scyna.Context
}

func NewAccountService(context *scyna.Context) *accountService {
	return &accountService{
		context:    context,
		Repository: &repository.AccountRepository{LOG: context.Logger},
	}
}
