package repository

import "github.com/go-rest-api/domain"

type CustomerRepository interface {
	CreateCustomer(*domain.Customer) (*domain.Customer, error)
	EditCustomer(*domain.Customer) error
	DeleteCustomer(id int64) error
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id int64) (*domain.Customer, error)
}
