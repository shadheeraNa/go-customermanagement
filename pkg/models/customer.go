package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shadheeraNa/go-customermanagement/pkg/config"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	MobileNum string `json:"MobileNum"`
	Orders    []Order
}

//type Order struct {
//	gorm.Model
//	OrderAddress string `json:"OrderAddress"`
//	OrderPrice   string `json:"OrderPrice"`
//	CustomerID   uint   `json:"CustomerID"`
//}

func init() {
	config.Connect()
	db = config.GetDB()
	//db.AutoMigrate(&Customer{}, &Order{})
	db.AutoMigrate(&Customer{}, &Order{})
}

func (c *Customer) CreateCustomer() *Customer {
	db.NewRecord(c) //new record comes with the gorm and it creates a new record for the order
	db.Create(&c)
	return c
}

func GetAllCustomers() []Customer {
	var Customers []Customer
	db.Find(&Customers)
	return Customers
}

func GetCustomerById(ID int64) (*Customer, *gorm.DB) {
	var getCustomer Customer
	db := db.Where("ID = ?", ID).Find(&getCustomer)
	return &getCustomer, db
}

func DeleteCustomer(ID int64) Customer {
	var customer Customer
	db.Where("ID = ?", ID).Delete(customer)
	return customer
}

func GetAll(db *gorm.DB) ([]Customer, error) {
	var customers []Customer
	err := db.Model(&Customer{}).Preload("Orders").Find(&customers).Error
	return customers, err
}
