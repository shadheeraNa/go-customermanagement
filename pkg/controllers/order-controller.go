package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shadheeraNa/go-customermanagement/pkg/models"
	"github.com/shadheeraNa/go-customermanagement/pkg/utils"
)

var NewOrder models.Order

// var NewOrder2 models.Customer

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	CreateOrder := &models.Order{}
	utils.ParseBody(r, CreateOrder)
	b := CreateOrder.CreateOrder()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	newOrders := models.GetAllOrders()
	res, _ := json.Marshal(newOrders)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	orderDetails, _ := models.GetOrderById(ID)
	res, _ := json.Marshal(orderDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var updateOrder = &models.Order{}
	utils.ParseBody(r, updateOrder)
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	orderDetails, db := models.GetOrderById(ID)
	// if updateOrder.OrderId != 0 {
	// 	orderDetails.OrderId = updateOrder.OrderId
	// }
	if updateOrder.OrderPrice != "" {
		orderDetails.OrderPrice = updateOrder.OrderPrice
	}
	if updateOrder.OrderAddress != "" {
		orderDetails.OrderAddress = updateOrder.OrderAddress
	}
	// if updateOrder.CustomerID != 0 {
	// 	orderDetails.CustomerID = updateOrder.CustomerID

	// }
	// if updateOrder.Customer.FirstName != "" {
	// 	orderDetails.Customer = updateOrder.Customer
	// }
	// if updateOrder.Customer.LastName != "" {
	// 	orderDetails.Customer = updateOrder.Customer
	// }
	// if updateOrder.Customer.MobileNum != "" {
	// 	orderDetails.Customer = updateOrder.Customer
	// }
	db.Save(&orderDetails)
	res, _ := json.Marshal(orderDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	ID, err := strconv.ParseInt(orderId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	order := models.DeleteOrder(ID)
	res, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
