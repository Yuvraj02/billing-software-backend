package router

import (
	"net/http"
	"restapi/billing-backend/internal/api/handlers"
)

func Router() *http.ServeMux{

	mux := http.NewServeMux()

	//All GET routes
	mux.HandleFunc("GET /", handlers.GetCustomers)
	mux.HandleFunc("GET /{id}", handlers.GetCustomerByID)
	mux.HandleFunc("GET /dim/{id}", handlers.GetDimensionsByID)
	
	
	//All POST routes
	mux.HandleFunc("POST /add_customer", handlers.AddCustomer)
	mux.HandleFunc("POST /add_dim", handlers.AddDimension)

	return mux 
}