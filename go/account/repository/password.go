package repository

import (
	"ex/account/model"

	scyna "github.com/scyna/core"
)

func (r *AccountRepository) LoadPassword(account *model.Account) scyna.Error {
	return nil
}

func (r *AccountRepository) UpdatePassword(cmd *scyna.Command, account *model.Account) scyna.Error {
	return nil
}
