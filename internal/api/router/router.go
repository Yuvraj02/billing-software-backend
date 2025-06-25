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
	mux.HandleFunc("GET /work_list", handlers.GetPendingWork)
	mux.HandleFunc("GET /work_list/{phone}", handlers.GetPendingWorkByPhone)
	mux.HandleFunc("GET /work_list/customer/{work_id}", handlers.GetPendingWorkById)
	mux.HandleFunc("GET /work_list/completed",handlers.GetCompletedWork)
	mux.HandleFunc("GET /work_list/completed/name/{name}",handlers.GetCompletedWorkByName)
	mux.HandleFunc("GET /work_list/completed/phone/{phone}", handlers.GetCompletedWorkByPhone)
	//GET : http://sanjhikala.com/work_list/name/yuvraj -->URL
		
	mux.HandleFunc("GET /work_list/name/{name}", handlers.GetPendingWorkByName)
	
	//All POST routes
	mux.HandleFunc("POST /add_customer", handlers.AddCustomer)
	mux.HandleFunc("POST /add_dim", handlers.AddDimension)
	mux.HandleFunc("POST /add_category", handlers.AddCategory)
	mux.HandleFunc("POST /add_work", handlers.AddWork)
	//All Patch routes
	mux.HandleFunc("PATCH /update_dim/{id}", handlers.PatchDimensions)
	mux.HandleFunc("PATCH /mark_complete/{id}", handlers.PatchWork)
	return mux 
}