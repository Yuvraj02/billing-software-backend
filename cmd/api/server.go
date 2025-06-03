package main

import (
	"fmt"
	"net/http"
	"os"
	"restapi/billing-backend/internal/api/router"
	"restapi/billing-backend/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	// fmt.Println("Working")
	err := godotenv.Load()
	if err!= nil {
		fmt.Println("Error in Loading .ENV File")
		return
	}

	err = sqlconnect.ConnectDB()
	if err!=nil{
		fmt.Printf("%s", err)
		return
	}

	port := os.Getenv("API_PORT")

	// http.HandleFunc("GET /", rootHandler)
	// http.HandleFunc("GET /{id}", userHandler)

	fmt.Println("Server Running on port", port)

	routerMux := router.Router()

	http.ListenAndServe(port,routerMux)
}
