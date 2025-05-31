package router

import (
	"net/http"
	"restapi/billing-backend/internal/api/handlers"
)

func Router() *http.ServeMux{

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetCustomers)

	return mux 
}