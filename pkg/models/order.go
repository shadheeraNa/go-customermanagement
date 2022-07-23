package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderAddress string `json:"OrderAddress"`
	OrderPrice   string `json:"OrderPrice"`
	CustomerID   uint   `json:"CustomerID"`
}

func (b *Order) CreateOrder() *Order {
	db.NewRecord(b) //new record comes with the gorm and it creates a new record for the order
	db.Create(&b)
	return b
}

func GetAllOrders() []Order {
	var Orders []Order
	db.Find(&Orders)
	return Orders
}

func GetOrderById(Id int64) (*Order, *gorm.DB) {
	var getOrder Order
	db := db.Where("ID = ?", Id).Find(&getOrder)
	return &getOrder, db
}

func DeleteOrder(ID int64) Order {
	var order Order
	db.Where("ID = ?", ID).Delete(order)
	return order
}
