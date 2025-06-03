//TODO : PATCH TO BE ADDED

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restapi/billing-backend/internal/models"
	"restapi/billing-backend/internal/repository/sqlconnect"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func GetDimensionsByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Fatal("Problem in String to Integer conversion")
	}

	var current_dim models.Dimension
	query := `SELECT customer_id, shoulder, upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, length, bottom FROM dimensions WHERE customer_id = $1`
	row := sqlconnect.Dbpool.QueryRow(context.Background(), query, id)
	err = row.Scan(&current_dim.CustomerId, &current_dim.Shoulder, &current_dim.UpperChest, &current_dim.Chest, &current_dim.Waist, &current_dim.Hip, &current_dim.Sleeves, &current_dim.NeckFront, &current_dim.NeckBack, &current_dim.Armhole, &current_dim.Length, &current_dim.Bottom)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "No Row found with id provided", http.StatusNotFound)
			return
		}
	}

	resonse := struct {
		Status string           `json:"status"`
		Data   models.Dimension `json:"data"`
	}{
		Status: "success",
		Data:   current_dim,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resonse)
}

func AddDimension(w http.ResponseWriter, r *http.Request) {

	var newDimensions models.Dimension
	json.NewDecoder(r.Body).Decode(&newDimensions)

	var newRow bool = false //This will check if the dimensions to be added has to be added as new record or update
	query := `SELECT customer_id FROM dimensions WHERE customer_id = $1`

	row := sqlconnect.Dbpool.QueryRow(context.Background(), query, newDimensions.CustomerId)
	err := row.Scan(&newDimensions.CustomerId)

	if err != nil {
		if err == pgx.ErrNoRows { //This means that data regarding current customer does not exist, so create a new row
			query = `INSERT INTO dimensions (customer_id, shoulder, upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, length, bottom) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
			_, err = sqlconnect.Dbpool.Exec(context.Background(), query, &newDimensions.CustomerId, &newDimensions.Shoulder, &newDimensions.UpperChest, &newDimensions.Chest, &newDimensions.Waist, &newDimensions.Hip, &newDimensions.Sleeves, &newDimensions.NeckFront, &newDimensions.NeckBack, &newDimensions.Armhole, &newDimensions.Length, &newDimensions.Bottom)
			newRow = true
		} else {
			http.Error(w, fmt.Sprintf("%s/n", err), http.StatusInternalServerError)
			return
		}
	}

	//If no new row is created then simply update the values into existing one
	if !newRow {
		query = `UPDATE dimensions SET customer_id = $1,shoulder=$2, upper_chest=$3, chest=$4, waist=$5, hip=$6, sleeves=$7, neck_front=$8, neck_back=$9, armhole=$10, length=$11, bottom=$12 WHERE customer_id = $13`
		_, err = sqlconnect.Dbpool.Exec(context.Background(), query, &newDimensions.CustomerId, &newDimensions.Shoulder, &newDimensions.UpperChest, &newDimensions.Chest, &newDimensions.Waist, &newDimensions.Hip, &newDimensions.Sleeves, &newDimensions.NeckFront, &newDimensions.NeckBack, &newDimensions.Armhole, &newDimensions.Length, &newDimensions.Bottom, &newDimensions.CustomerId)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}
	}

	response := struct {
		Status string           `json:"status"`
		Data   models.Dimension `json:"data"`
	}{
		Status: "success",
		Data:   newDimensions,
	}

	newRow = false
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
