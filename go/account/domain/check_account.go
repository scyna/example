package domain

import (
	"ex/account/model"

	scyna "github.com/scyna/core"
)

func (account *accountService) AssureAccountNotExists(email model.EmailAddress) scyna.Error {
	if _, err := account.Repository.GetAccountByEmail(email); err == nil {
		return model.ACCOUNT_EXISTED
	}
	return nil
}
