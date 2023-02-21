package repository

import (
	"ex/vertical_slide/shared/model"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *repository) CreateAccount(account *model.Account) scyna.Error {

	if err := qb.Insert(ACCOUNT_TABLE).
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(account).
		ExecRelease(); err != nil {
		r.context.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
}
