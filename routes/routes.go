package routes

import (
	"github.com/go-rest-api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/go-rest-api/controllers"
	"github.com/gorilla/handlers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/customers", controllers.GetAllCustomers).Methods("GET")
	r.HandleFunc("/api/customers/{id}", controllers.GetCustomer).Methods("GET")
	r.HandleFunc("/api/customers", controllers.CreateCustomer).Methods("POST")
	r.HandleFunc("/api/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	r.HandleFunc("/api/customers/{id}", controllers.EditCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9005", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
