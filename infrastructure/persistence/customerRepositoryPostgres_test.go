package persistence

import (
	"github.com/go-rest-api/domain"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestCustomerRepositoryPostgres_CreateCustomer(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		customer *domain.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomerRepositoryPostgres{
				DB: tt.fields.DB,
			}
			if err := r.CreateCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomerRepositoryPostgres_DeleteCustomer(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomerRepositoryPostgres{
				DB: tt.fields.DB,
			}
			if err := r.DeleteCustomer(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomerRepositoryPostgres_EditCustomer(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		customer *domain.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomerRepositoryPostgres{
				DB: tt.fields.DB,
			}
			if err := r.EditCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("EditCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomerRepositoryPostgres_GetAllCustomers(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomerRepositoryPostgres{
				DB: tt.fields.DB,
			}
			got, err := r.GetAllCustomers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerRepositoryPostgres_GetCustomer(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomerRepositoryPostgres{
				DB: tt.fields.DB,
			}
			got, err := r.GetCustomer(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomerRepositoryPostgres(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *CustomerRepositoryPostgres
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerRepositoryPostgres(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerRepositoryPostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
