package change_password

import (
	"ex/vertical_slide/shared/proto"

	scyna "github.com/scyna/core"
)

const URL = "/account/create"

func Handler(ctx *scyna.Endpoint, request *proto.ChangePasswordRequest) scyna.Error {
	ctx.Info("Receive ChangePasswordRequest")

	/*TODO*/

	return nil
}
