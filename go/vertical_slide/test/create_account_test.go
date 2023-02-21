package test

import (
	"ex/vertical_slide/features/create_account"
	"ex/vertical_slide/shared/model"
	"ex/vertical_slide/shared/proto"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(create_account.URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)

	scyna_test.EndpointTest(create_account.URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectError(model.ACCOUNT_EXISTED).Run(t)
}

func TestCreateBadEmail(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(create_account.URL).
		WithRequest(&proto.CreateAccountRequest{
			Email: "a+gmail.com",
			Name:  "Nguyen Van A",
		}).
		ExpectError(scyna.REQUEST_INVALID).Run(t)
}
