package repository

import scyna "github.com/scyna/core"

const TABLE_NAME = "ex_account.account"

type repository struct {
	context *scyna.Context
}

func LoadRepository(context *scyna.Context) *repository {
	return &repository{context: context}
}
