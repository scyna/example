package repository

import (
	"ex/customer/domain"

	scyna "github.com/scyna/core"
)

const CUSTOMER_TABLE = "ex_customer.customer"

type customerRepository struct {
	context scyna.Context
}

func NewRepository(context scyna.Context) domain.IRepository {
	return &customerRepository{context: context}
}
