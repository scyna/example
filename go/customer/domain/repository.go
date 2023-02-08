package domain

import (
	"ex/customer/model"

	scyna "github.com/scyna/core"
)

var OneID = scyna.InitSerialNumber("ddd.customer.oneid")

type IRepository interface {
	CreateCustomerProfile(customer *model.Customer) scyna.Error
	/*TODO: methods*/
}

type RepositoryCreator func(LOG scyna.Logger) IRepository

var repositoryCreator RepositoryCreator

func LoadRepository(LOG scyna.Logger) IRepository {
	if repositoryCreator == nil {
		panic("No RepositoryCreator attached")
	}
	return repositoryCreator(LOG)
}

func AttachRepositoryCreator(rc RepositoryCreator) {
	repositoryCreator = rc
}
