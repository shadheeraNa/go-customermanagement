package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shadheeraNa/go-customermanagement/pkg/config"
	"github.com/shadheeraNa/go-customermanagement/pkg/models"
	"github.com/shadheeraNa/go-customermanagement/pkg/utils"
	"net/http"
	"strconv"
)

var NewCustomer models.Customer

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	CreateCustomer := &models.Customer{}
	utils.ParseBody(r, CreateCustomer)
	b := CreateCustomer.CreateCustomer()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	newCustomers, _ := models.GetAll(config.GetDB())
	//newCustomers := models.GetAllCustomers()
	res, _ := json.Marshal(newCustomers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	customerDetails, _ := models.GetCustomerById(ID)
	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var updateCustomer = &models.Customer{}
	utils.ParseBody(r, updateCustomer)
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	customerDetails, db := models.GetCustomerById(ID)
	// if updateCustomer.CustomerID != 0 {
	// 	customerDetails.CustomerID = updateCustomer.CustomerID
	// }
	if updateCustomer.FirstName != "" {
		customerDetails.FirstName = updateCustomer.FirstName
	}
	if updateCustomer.LastName != "" {
		customerDetails.LastName = updateCustomer.LastName
	}
	if updateCustomer.MobileNum != "" {
		customerDetails.MobileNum = updateCustomer.MobileNum
	}
	db.Save(&customerDetails)
	res, _ := json.Marshal(customerDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	ID, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	customer := models.DeleteCustomer(ID)
	res, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
