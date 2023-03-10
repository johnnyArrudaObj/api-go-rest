package persistence

import (
	"github.com/go-rest-api/domain"
	"github.com/go-rest-api/infrastructure/database"
	"gorm.io/gorm"
)

type CustomerRepositoryPostgres struct {
	DB *gorm.DB
}

func NewCustomerRepositoryPostgres(db *gorm.DB) *CustomerRepositoryPostgres {
	return &CustomerRepositoryPostgres{DB: db}
}

func (r *CustomerRepositoryPostgres) CreateCustomer(customer *domain.Customer) error {
	if result := database.DB.Create(&customer); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerRepositoryPostgres) EditCustomer(customer *domain.Customer) error {
	if result := database.DB.Save(&customer); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerRepositoryPostgres) DeleteCustomer(id int64) error {
	var customer domain.Customer
	if result := database.DB.Delete(&customer, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerRepositoryPostgres) GetAllCustomers() ([]domain.Customer, error) {
	var p []domain.Customer
	if result := database.DB.Find(&p); result.Error != nil {
		return []domain.Customer{}, result.Error
	}
	return p, nil
}

func (r *CustomerRepositoryPostgres) GetCustomer(id int64) (*domain.Customer, error) {
	var customer domain.Customer
	if result := database.DB.First(&customer, id); result.Error != nil {
		return &domain.Customer{}, result.Error
	}
	return &customer, nil
}
