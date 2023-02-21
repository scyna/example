package repository

import scyna "github.com/scyna/core"

const ACCOUNT_TABLE = "ex_account.account"

type BaseRepository struct {
	LOG *scyna.Logger
}

func NewBaseRepository(LOG *scyna.Logger) *BaseRepository {
	return &BaseRepository{LOG: LOG}
}
