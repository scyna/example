package repository

import scyna "github.com/scyna/core"

const ACCOUNT_TABLE = "ex_account.account"

type repository struct {
	context *scyna.Context
}

func LoadRepository(context *scyna.Context) *repository {
	return &repository{context: context}
}
