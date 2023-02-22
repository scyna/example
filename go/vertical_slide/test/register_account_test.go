package test

import (
	"ex/vertical_slide/features/register_account"
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestRegisterAccount_Success(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(register_account.PATH).
		WithRequest(&proto.RegisterAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)
}

func TestRegisterAccount_Duplicated(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(register_account.PATH).
		WithRequest(&proto.RegisterAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)

	scyna_test.EndpointTest(register_account.PATH).
		WithRequest(&proto.RegisterAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectError(model.ACCOUNT_EXISTED).Run(t)
}

func TestRegisterAccount_BadEmail(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(register_account.PATH).
		WithRequest(&proto.RegisterAccountRequest{
			Email: "a+gmail.com",
			Name:  "Nguyen Van A",
		}).
		ExpectError(scyna.REQUEST_INVALID).Run(t)
}
