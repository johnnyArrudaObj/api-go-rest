package domain

type Customer struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
	Id   int    `json:"id"`
}

var Customers []Customer
