package main

import (
	"fmt"
	"net/http"
	"restapi/billing-backend/internal/api/router"
)


// func userHandler(w http.ResponseWriter, r *http.Request){
	
// 	idStr := r.PathValue("id")
// 	id,err := strconv.Atoi(idStr)
// 	if err != nil{
// 		fmt.Println("Error in converting id to int")
// 		return
// 	}

// 	response := struct {
// 		Status string `json:"status"`
// 		Data models.User `json:"data"`
// 	}{
// 		Status: "success",
// 		Data: initialMap[id],
// 		}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)

// }

func main() {

	// fmt.Println("Working")
	port := ":3000"


	// http.HandleFunc("GET /", rootHandler)
	// http.HandleFunc("GET /{id}", userHandler)

	fmt.Println("Server Running on port", port)

	routerMux := router.Router()

	http.ListenAndServe(port,routerMux)
}
