package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/billing-backend/internal/models"
)

var (
	customersMap = make(map[int]models.Customer)
	next         = 0
)

func init() {

	customersMap[0] = models.Customer{Id: 0, Name: "Yuvraj", Email: nil, Phone: "877080213", UserID: 0, Address: nil}
	next++

}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	var customersList []models.Customer
	for _, value := range customersMap {
		customersList = append(customersList, value)
	}

	response := struct {
		Status string            `json:"status"`
		Count  int               `json:"customers_count"`
		Data   []models.Customer `json:"data"`
	}{
		Status: "success",
		Count:  len(customersList),
		Data:   customersList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
