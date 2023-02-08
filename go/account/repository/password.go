package repository

import (
	"ex/account/model"

	scyna "github.com/scyna/core"
)

func (r *accountRepository) LoadPassword(account *model.Account) scyna.Error {
	return nil
}

func (r *accountRepository) UpdatePassword(cmd *scyna.Command, account *model.Account) scyna.Error {
	return nil
}
