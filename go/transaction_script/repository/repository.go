package repository

import scyna "github.com/scyna/core"

const TABLE_NAME = "ex_account.account"

type repository struct {
	LOG scyna.Logger
}

func LoadRepository(LOG scyna.Logger) *repository {
	return &repository{LOG: LOG}
}
