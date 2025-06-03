package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"restapi/billing-backend/internal/models"
	"restapi/billing-backend/internal/repository/sqlconnect"

	"github.com/jackc/pgx/v5"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	query := `SELECT category_id, category_name, shoulder, upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, length, bottom FROM categories`

	rows, _ := sqlconnect.Dbpool.Query(context.Background(), query)

	categoryList, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Category])
	if err != nil {
		http.Error(w, "Problem in Getting categories", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string             `json:"status"`
		Count  int                `json:"category_count"`
		Data   []*models.Category `json:"category_data"`
	}{
		Status: "success",
		Count:  len(categoryList),
		Data:   categoryList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func AddCategory(w http.ResponseWriter, r *http.Request) {

	var newCategory models.Category
	json.NewDecoder(r.Body).Decode(&newCategory)

	query := `INSERT INTO categories VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := sqlconnect.Dbpool.Exec(context.Background(), query, &newCategory.CategoryName, &newCategory.Shoulder, &newCategory.UpperChest, &newCategory.Chest, &newCategory.Waist, &newCategory.Hip, &newCategory.Sleeves, &newCategory.NeckFront, &newCategory.NeckBack, &newCategory.Armhole, &newCategory.Length, &newCategory.Bottom)
	if err!=nil{
		http.Error(w,"Error in Writing New Category to Database", http.StatusInternalServerError)
		return
	}

	response:= struct{
		Status string `json:"status"`
		Data models.Category `json:"data"`
	}{
		Status: "success",
		Data: newCategory,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	
}
