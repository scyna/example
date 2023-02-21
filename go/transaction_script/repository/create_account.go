package repository

import (
	"ex/transaction_script/model"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *repository) CreateAccount(account *model.Account) scyna.Error {

	if err := qb.Insert(TABLE_NAME).
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(account).
		ExecRelease(); err != nil {
		r.LOG.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
}
