package routes

import (
	"github.com/gorilla/mux"
	"github.com/shadheeraNa/go-customermanagement/pkg/controllers"
)

var RegisterCustomerManagementRoutes = func(router *mux.Router) { //this function create handlers
	router.HandleFunc("/customer/", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{customerId}", controllers.DeleteCustomer).Methods("DELETE")
}
