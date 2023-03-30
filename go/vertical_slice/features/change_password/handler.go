package change_password

import (
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"

	scyna "github.com/scyna/core"
)

const PATH = "/account/change-password"

func Handler(ctx *scyna.Endpoint, request *proto.ChangePasswordRequest) scyna.Error {
	ctx.Info("Receive ChangePasswordRequest")

	repository := NewRepository(ctx)

	account, err := repository.GetAccountByID(request.ID)

	if err != nil {
		return err
	}

	if account.Password != request.Current {
		return model.PASSWORD_NOT_MATCH
	}

	account.Password = request.Future

	if err := repository.ChangePassword(account); err != nil {
		return err
	}

	return nil
}
