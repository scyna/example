package test

import (
	"ex/transaction_script/service"
	"os"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init("scyna_test")

	/*register services*/
	scyna.RegisterEndpoint(service.CREATE_ACCOUNT_URL, service.CreateAccountHandler)
	scyna.RegisterEndpoint(service.GET_ACCOUNT_URL, service.GetAccountHandler)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE ex_account.account", nil).ExecRelease()
}
