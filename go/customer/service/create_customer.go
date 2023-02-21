package service

import (
	"ex/customer/domain"
	"ex/customer/model"
	proto "ex/customer/proto"

	scyna "github.com/scyna/core"
)

func CreateCustomerHandler(ctx *scyna.Endpoint, request *proto.CreateCustomerRequest) scyna.Error {
	var ret scyna.Error
	repository := domain.LoadRepository(ctx)
	customer := model.Customer{ID: model.CustomerID(domain.OneID.Next())}

	if customer.Identity, ret = model.NewIdentity(request.IDType, request.IDNumber); ret != nil {
		return ret
	}

	if ret = domain.IdentityExists(repository, customer.Identity); ret != nil {
		return ret
	}

	if ret := repository.CreateCustomerProfile(&customer); ret != nil {
		return ret
	}

	ctx.Response(&proto.CreateCustomerResponse{OneID: string(customer.ID)})

	return scyna.OK
}
