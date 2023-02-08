package model

import scyna "github.com/scyna/core"

type Name string
type CustomerID string

type Customer struct {
	Context  scyna.Context
	ID       CustomerID
	Identity Identity
	Name     CustomerName
	Gender   Gender
	Email    EmailAddress
	Phone    PhoneNumber
	Address  StreetAddress
	DOB      DateOfBirth
}
