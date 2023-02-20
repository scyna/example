package main

import (
	"vertical_slide/account/features/create_account"

	scyna_setup "github.com/scyna/core/setup"
)

func main() {
	scyna_setup.Init()
	scyna_setup.NewModule("ex_account", "123456").
		AddConsumer("ex_account").
		Build()

	scyna_setup.NewClient("account_test", "123456").
		UseEndpoint(create_account.URL).
		//UseEndpoint(service.GET_ACCOUNT_URL).
		Build()

}
