package service

import (
	"ex/hello/proto"

	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
)

func HelloHandler(ctx scyna.Context, request *proto.HelloRequest) scyna.Error {
	ctx.Info("Receive HelloRequest")

	if err := validateHelloRequest(request); err != nil {
		return scyna.REQUEST_INVALID
	}

	return ctx.OK(&proto.HelloResponse{Content: "Hello " + request.Name})
}

func validateHelloRequest(request *proto.HelloRequest) error {
	return validation.Validate(request.Name, validation.Required, validation.Length(3, 40))
}
