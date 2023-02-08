package repository

import (
	"ex/account/domain"

	scyna "github.com/scyna/core"
)

const ACCOUNT_TABLE = "ex_account.account"

type accountRepository struct {
	LOG scyna.Logger
}

func NewRepository(LOG scyna.Logger) domain.IRepository {
	return &accountRepository{LOG: LOG}
}
