package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *repository) GetAccount(email string) (*Account, scyna.Error) {
	var ret Account
	if err := qb.Select(TABLE_NAME).
		Columns("id", "name", "email").
		Where(qb.Eq("email")).
		Query(scyna.DB).
		Bind(email).
		GetRelease(&ret); err != nil {
		r.context.Error(err.Error())
		return nil, ACCOUNT_NOT_EXISTED
	}
	return &ret, nil
}
