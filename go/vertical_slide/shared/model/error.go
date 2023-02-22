package model

import (
	scyna "github.com/scyna/core"
)

var (
	ACCOUNT_EXISTED     = scyna.NewError(100, "User Existed")
	ACCOUNT_NOT_EXISTED = scyna.NewError(101, "User Not Existed")
	PASSWORD_NOT_MATCH  = scyna.NewError(102, "Password Not Match")
)
