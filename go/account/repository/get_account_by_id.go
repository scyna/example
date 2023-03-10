package repository

import (
	"ex/account/model"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func (r *AccountRepository) GetAccountByID(ID uint64) (*model.Account, scyna.Error) {

	var account struct {
		ID    uint64 `db:"id"`
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	if err := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).Bind(ID).GetRelease(&account); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.ACCOUNT_NOT_EXISTED
	}

	ret := &model.Account{
		LOG:  r.LOG,
		ID:   account.ID,
		Name: account.Name,
	}

	var err scyna.Error
	if ret.Email, err = model.ParseEmail(account.Email); err != nil {
		return nil, model.BAD_EMAIL
	}

	return ret, nil
}
