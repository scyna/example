package domain

import (
	"ex/account/model"

	scyna "github.com/scyna/core"
)

func AssureAccountNotExists(repository IRepository, email model.EmailAddress) scyna.Error {
	if _, err := repository.GetAccountByEmail(email); err == nil {
		return model.USER_EXISTED
	}
	return nil
}
