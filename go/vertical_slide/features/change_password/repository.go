package change_password

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

func (r *Repository) ChangePassword(account *model.Account) scyna.Error {

	if applied, _ := qb.Update(repository.ACCOUNT_TABLE).
		Set("password").
		Where(qb.Eq("id")).
		Existing().
		Query(scyna.DB).
		Bind(account.Password, account.ID).
		ExecCASRelease(); !applied {
		return scyna.SERVER_ERROR
	}

	return nil
}
