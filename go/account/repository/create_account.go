package repository

import (
	"ex/account/model"

	scyna "github.com/scyna/core"
)

const _QUERY = "INSERT INTO " + ACCOUNT_TABLE + "(id, name, email, password) VALUES(?,?,?,?)"

func (r *accountRepository) CreateAccount(cmd *scyna.Command, account *model.Account) scyna.Error {
	cmd.Batch().Query(_QUERY, account.ID, account.Name, account.Email.String(), account.Password.String())
	return nil
}
