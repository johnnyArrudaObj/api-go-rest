package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-rest-api/domain"
	"github.com/go-rest-api/infrastructure/database"
	"github.com/go-rest-api/infrastructure/persistence"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	repo := persistence.NewCustomerRepositoryPostgres(database.DB)
	customers, err := repo.GetAllCustomers()
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := persistence.NewCustomerRepositoryPostgres(database.DB)
	customer, err := repo.GetCustomer(int64(idInt))
	if err != nil {
		log.Fatal(err.Error())
	}
	json.NewEncoder(w).Encode(customer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	repo := persistence.NewCustomerRepositoryPostgres(database.DB)
	err := repo.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	var customer domain.Customer
	database.DB.Delete(&customer, idInt)
	json.NewEncoder(w).Encode(customer)
}

func EditCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := persistence.NewCustomerRepositoryPostgres(database.DB)
	customer, err := repo.GetCustomer(int64(idInt))
	if err != nil {
		return
	}

	json.NewDecoder(r.Body).Decode(customer)
	err = repo.EditCustomer(customer)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(customer)
}
