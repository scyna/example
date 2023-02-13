package repository

import (
	scyna "github.com/scyna/core"
)

const ACCOUNT_TABLE = "ex_account.account"

type AccountRepository struct {
	LOG scyna.Logger
}
