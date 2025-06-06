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
	mux.HandleFunc("GET /customers/{phone}",handlers.GetCustomerByPhone)
	mux.HandleFunc("GET /dim/{phone}", handlers.GetDimensionsByPhone)
	mux.HandleFunc("GET /categories", handlers.GetCategories)
	
	//All POST routes
	mux.HandleFunc("POST /add_customer", handlers.AddCustomer)
	mux.HandleFunc("POST /add_dim", handlers.AddDimension)
	mux.HandleFunc("POST /add_category", handlers.AddCategory)

	//All Patch routes
	mux.HandleFunc("PATCH /update_dim/{id}", handlers.PatchDimensions)
	return mux 
}