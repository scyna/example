package repository

import (
	"ex/vertical_slide/shared/model"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *BaseRepository) GetAccountByID(ID uint64) (*model.Account, scyna.Error) {
	var ret model.Account
	if err := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(ID).
		GetRelease(&ret); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.ACCOUNT_NOT_EXISTED
	}
	return &ret, nil
}
