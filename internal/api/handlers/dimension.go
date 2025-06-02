package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/billing-backend/internal/models"
	"strconv"
)

var (
	dimensionsMap = make(map[int]models.Dimension)
	nextDim         = 0
)

//This helper function will take a literal / constant float value, store it and return the address of memory area, where it was stored
//Doing this will help in storing values as pointers in our dimension model, which essentially has null values depending upon the product that a customer chooses
func float32Ptr(f float32) *float32{
	return &f
}

func init() {

	dimensionsMap[0] = models.Dimension{CustomerId: 0, Length: float32Ptr(32.5), NeckFront: float32Ptr(22.5), UpperChest: float32Ptr(12.5), Chest: float32Ptr(10), Waist: float32Ptr(22), Hip: float32Ptr(28.5)}
	nextDim++
	dimensionsMap[1] = models.Dimension{CustomerId: 1, Length: float32Ptr(32.5),Bottom: float32Ptr(20) }
	nextDim++

}

func GetDimensionsByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id ,err := strconv.Atoi(idStr)
	
	if err != nil {
		log.Fatal("Problem in String to Integer conversion")
	}

	data := dimensionsMap[id]

	response := struct {
		Status string `json:"status"`
		Data models.Dimension
	}{
		Status: "success",
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func AddDimension(w http.ResponseWriter, r *http.Request){

	var newDimensions models.Dimension
	json.NewDecoder(r.Body).Decode(&newDimensions)


	dimensionsMap[newDimensions.CustomerId] = newDimensions

	w.WriteHeader(http.StatusAccepted)

}