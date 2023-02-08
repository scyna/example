package test

import (
	"ex/transaction_script/proto"
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
}

func TestCreateThenGet(t *testing.T) {
	cleanup()
	var response proto.CreateAccountResponse
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t, &response)

	scyna_test.EndpointTest(service.GET_ACCOUNT_URL).
		WithRequest(&proto.GetAccountByEmailRequest{Email: "a@gmail.com"}).
		ExpectResponse(&proto.GetAccountResponse{
			Id:    response.Id,
			Email: "a@gmail.com",
			Name:  "Nguyen Van A",
		}).Run(t)
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
