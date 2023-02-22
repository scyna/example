package test

import (
	"ex/vertical_slide/features/get_account"
	"ex/vertical_slide/features/register_account"
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"
	"testing"

	scyna_test "github.com/scyna/core/testing"
)

func TestGetAccount_Success(t *testing.T) {
	cleanup()

	var response proto.RegisterAccountResponse
	scyna_test.EndpointTest(register_account.PATH).
		WithRequest(&proto.RegisterAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t, &response)

	scyna_test.EndpointTest(get_account.PATH).
		WithRequest(&proto.GetAccountByEmailRequest{
			Email: "a@gmail.com",
		}).
		ExpectResponse(&proto.GetAccountResponse{
			ID:    response.ID,
			Email: "a@gmail.com",
			Name:  "Nguyen Van A",
		}).Run(t)
}

func TestGetAccount_NotExists(t *testing.T) {
	cleanup()

	scyna_test.EndpointTest(get_account.PATH).
		WithRequest(&proto.GetAccountByEmailRequest{
			Email: "a@gmail.com",
		}).
		ExpectError(model.ACCOUNT_NOT_EXISTED).Run(t)
}
