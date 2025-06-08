package main

import (
	"fmt"
	"net/http"
	"os"
	"restapi/billing-backend/internal/api/middlewares"
	"restapi/billing-backend/internal/api/router"
	"restapi/billing-backend/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {


	err := godotenv.Load()
	if err!= nil {
		fmt.Printf("%s\n",err)
		return
	}

	err = sqlconnect.ConnectDB()
	if err!=nil{
		fmt.Printf("%s", err)
		return
	}

	port := os.Getenv("API_PORT")

	fmt.Println("Server Running on port", port)

	routerMux := router.Router()
	cors := middlewares.Cors(routerMux)
	http.ListenAndServe(port,cors)
}
