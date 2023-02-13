package domain

import (
	scyna "github.com/scyna/core"
)

var (
	ACCOUNT_EXISTED     = scyna.NewError(100, "User Existed")
	ACCOUNT_NOT_EXISTED = scyna.NewError(101, "User NOT Existed")
)
