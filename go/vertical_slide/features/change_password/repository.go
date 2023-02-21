package change_password

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/repository"

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

func (r *Repository) ChangePasword(account *model.Account, password string) scyna.Error {

	return nil
}
