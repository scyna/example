package domain

import (
	"ex/account/model"
	"ex/account/repository"

	scyna "github.com/scyna/core"
)

type accountService struct {
	Repository *repository.AccountRepository
	context    scyna.Context
}

func NewAccountService(context scyna.Context) *accountService {
	return &accountService{
		context:    context,
		Repository: &repository.AccountRepository{LOG: context},
	}
}

func (account *accountService) AssureAccountNotExists(email model.EmailAddress) scyna.Error {
	if _, err := account.Repository.GetAccountByEmail(email); err == nil {
		return model.ACCOUNT_EXISTED
	}
	return nil
}
