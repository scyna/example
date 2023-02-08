package repository

import (
	"ex/customer/model"

	scyna "github.com/scyna/core"
)

func (r *customerRepository) GetCustomerByIdentity(identity model.Identity) (*model.Customer, scyna.Error) {
	/*TODO*/
	return &model.Customer{}, nil
}
