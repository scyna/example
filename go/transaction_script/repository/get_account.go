package repository

import (
	"ex/transaction_script/model"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *repository) GetAccount(email string) (*model.Account, scyna.Error) {
	var ret model.Account
	if err := qb.Select(TABLE_NAME).
		Columns("id", "name", "email").
		Where(qb.Eq("email")).
		Query(scyna.DB).
		Bind(email).
		GetRelease(&ret); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.ACCOUNT_NOT_EXISTED
	}
	return &ret, nil
}
