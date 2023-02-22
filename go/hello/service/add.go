package service

import (
	"ex/hello/proto"

	scyna "github.com/scyna/core"
)

func AddHandler(ctx scyna.Context, request *proto.AddRequest) scyna.Error {
	ctx.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		return ADD_RESULT_TOO_BIG
	}

	return ctx.OK(&proto.AddResponse{Sum: sum})
}
