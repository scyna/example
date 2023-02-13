package test

import (
	"ex/transaction_script/domain"
	"ex/transaction_script/proto"
	"ex/transaction_script/service"
	"testing"

	scyna_test "github.com/scyna/core/testing"
)

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

func TestGetFail(t *testing.T) {
	cleanup()

	scyna_test.EndpointTest(service.GET_ACCOUNT_URL).
		WithRequest(&proto.GetAccountByEmailRequest{Email: "a@gmail.com"}).
		ExpectError(domain.ACCOUNT_NOT_EXISTED).
		Run(t)
}
