package create_account

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

const TABLE_NAME = "ex_account.account"

type repository struct {
	context *scyna.Context
}

func LoadRepository(context *scyna.Context) *repository {
	return &repository{context: context}
}

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

func (r *repository) CreateAccount(account *Account) scyna.Error {

	if err := qb.Insert(TABLE_NAME).
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(account).
		ExecRelease(); err != nil {
		r.context.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
}
