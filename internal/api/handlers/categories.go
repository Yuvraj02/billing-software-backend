package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/billing-backend/internal/models"
)

var (
	categoryMap = make(map[int]models.Category)
	catID       = 0
)

func StringPTR(s string) *string {
	return &s
}

func init() {

	categoryMap[catID] = models.Category{CategoryId: catID, CategoryName: "Kurta", Length: StringPTR("Length"), NeckFront: StringPTR("Neck Front"), UpperChest: StringPTR("Upper Chest"), Chest: StringPTR("Chest"), Waist: StringPTR("Waist"), Hip: StringPTR("Hip")}
	catID++
	categoryMap[catID] = models.Category{CategoryId: catID, CategoryName: "Salvaar", Length: StringPTR("Length"), Bottom: StringPTR("Bottom")}
	catID++

}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	var categoryList []models.Category

	for _, value := range categoryMap {
		categoryList = append(categoryList, value)
	}

	response := struct {
		Status string            `json:"status"`
		Count  int               `json:"category_count"`
		Data   []models.Category `json:"category_list"`
	}{
		Status: "success",
		Count : len(categoryList),
		Data: categoryList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func AddCategory(w http.ResponseWriter, r *http.Request){

	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Error in Parsing Category JSON", http.StatusBadRequest)
		return
	}

	newCategory.CategoryId = catID
	categoryMap[catID] = newCategory
	catID++
}
