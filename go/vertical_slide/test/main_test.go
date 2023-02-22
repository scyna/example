package test

import (
	"ex/vertical_slide/features/get_account"
	"ex/vertical_slide/features/register_account"
	"os"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init("scyna_test")

	/*register services*/
	scyna.RegisterEndpoint(register_account.PATH, register_account.Handler)
	scyna.RegisterEndpoint(get_account.PATH, get_account.Handler)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE ex_account.account", nil).ExecRelease()
}
