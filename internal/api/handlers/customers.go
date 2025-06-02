package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/billing-backend/internal/models"
	"strconv"
)

var (
	customersMap = make(map[int]models.Customer)
	next         = 0
)

func init() {

	customersMap[0] = models.Customer{Id: 0, Name: "Yuvraj", Email: nil, Phone: "877080213", UserID: 0, Address: nil}
	next++
	customersMap[1] = models.Customer{Id: 1, Name: "John", Email: nil, Phone: "877080213", UserID: 0, Address: nil}
	next++

}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalln("Error Parsing String to int")
	}

	customerData := customersMap[id]

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customerData)

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

func AddCustomer(w http.ResponseWriter, r *http.Request) {

	var newCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)

	newCustomer.Id = next
	customersMap[next] = newCustomer
	next++

	response := struct {
		Status string          `json:"status"`
		Data   models.Customer `json:"customer_data"`
	}{
		Status: "success",
		Data:   newCustomer,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
