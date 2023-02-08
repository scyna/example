package test

import (
	"ex/transaction_script/proto"
	"ex/transaction_script/repository"
	"ex/transaction_script/service"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)

	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectError(repository.ACCOUNT_EXISTED).Run(t)

}

func TestCreateBadEmail(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.GetAccountResponse{
			Email: "a+gmail.com",
			Name:  "Nguyen Van A",
		}).
		ExpectError(scyna.REQUEST_INVALID).Run(t)
}
