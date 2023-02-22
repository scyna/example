package register_account

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/repository"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

type Repository struct {
	repository.BaseRepository
}

func NewRepository(LOG scyna.Logger) *Repository {
	return &Repository{
		BaseRepository: *repository.NewBaseRepository(LOG),
	}
}

func (r *Repository) CreateAccount(account *model.Account) scyna.Error {

	if err := qb.Insert(repository.ACCOUNT_TABLE).
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(account).
		ExecRelease(); err != nil {
		r.LOG.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
}
