//TODO : PATCH TO BE ADDED

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"restapi/billing-backend/internal/models"
	"restapi/billing-backend/internal/repository/sqlconnect"
	"strconv"
	"strings"
	"github.com/jackc/pgx/v5"
)

/*
	Workflow of Adding Dimensions :-
	->User will enter the mobile number in the client
	->Client will request the dimensions according to the phone number
	->Server when recieves the request, initiates a retrieval request for the dimension/s corresponding to the phone number from database
	->If there is a record existing, then a patch request will be sent from the client side with updated dimensions
	->else a new record is created and stored with the provided phone number
*/

func GetDimensionsByPhone(w http.ResponseWriter, r *http.Request) {

	phoneStr := r.PathValue("phone")

	query := `SELECT customer_id, customer_name, customer_phone, length,shoulder,upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, bottom FROM dimensions WHERE customer_phone = $1`

	rows, err := sqlconnect.Dbpool.Query(context.Background(), query, phoneStr)
	dimensionsList, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Dimension])

	// err = row.Scan(&current_dim.CustomerId, &current_dim.Shoulder, &current_dim.UpperChest, &current_dim.Chest, &current_dim.Waist, &current_dim.Hip, &current_dim.Sleeves, &current_dim.NeckFront, &current_dim.NeckBack, &current_dim.Armhole, &current_dim.Length, &current_dim.Bottom)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "No Row found with phone provided", http.StatusNotFound)
			return
		}
	}

	resonse := struct {
		Status string              `json:"status"`
		Data   []*models.Dimension `json:"data"`
	}{
		Status: "success",
		Data:   dimensionsList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resonse)
}

func AddDimension(w http.ResponseWriter, r *http.Request) {

	var newDimensions models.Dimension
	json.NewDecoder(r.Body).Decode(&newDimensions)

	query := `INSERT INTO dimensions (customer_id, customer_name, customer_phone, shoulder, upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, length, bottom) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`
	_, err := sqlconnect.Dbpool.Exec(context.Background(), query, &newDimensions.CustomerId, &newDimensions.CustomerName, &newDimensions.CustomerPhone, &newDimensions.Shoulder, &newDimensions.UpperChest, &newDimensions.Chest, &newDimensions.Waist, &newDimensions.Hip, &newDimensions.Sleeves, &newDimensions.NeckFront, &newDimensions.NeckBack, &newDimensions.Armhole, &newDimensions.Length, &newDimensions.Bottom)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s/n", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Data   models.Dimension `json:"data"`
	}{
		Status: "success",
		Data:   newDimensions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func PatchDimensions(w http.ResponseWriter, r *http.Request) {

	// var updateMap map[string]interface{}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Error in converting dimension Id ", http.StatusInternalServerError)
		return
	}

	queryParams := r.URL.Query()
	// fmt.Println("Path param is :", path, "Query Param is : ", queryParams)
	name := queryParams.Get("name")
	// fmt.Println("Path param is :", id, "Name is : ", name)
	var args []interface{}
	args = append(args, id)

	query := `SELECT customer_id, customer_name, customer_phone, length,shoulder,upper_chest, chest, waist, hip, sleeves, neck_front, neck_back, armhole, bottom FROM dimensions WHERE 1=1`
	query, args = addFilter(query, name, args)

	var updateMap map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateMap)
	if err != nil {
		http.Error(w, "Error Parsing JSON from Dimensions request", http.StatusInternalServerError)
		return
	}
	//fmt.Println(query)

	var existingDimension models.Dimension
	row := sqlconnect.Dbpool.QueryRow(context.Background(), query, args...)
	err = row.Scan(&existingDimension.CustomerId, &existingDimension.CustomerName, &existingDimension.CustomerPhone, &existingDimension.Length, &existingDimension.Shoulder, &existingDimension.UpperChest, &existingDimension.Chest, &existingDimension.Waist, &existingDimension.Hip, &existingDimension.Sleeves, &existingDimension.NeckFront, &existingDimension.NeckBack, &existingDimension.Armhole, &existingDimension.Bottom)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
		return
	}

	//Apply patch update to the existing
	dimensionVal := reflect.ValueOf(&existingDimension).Elem()
	dimensionType := dimensionVal.Type()

	for key, value := range updateMap {
		for i := 0; i < dimensionVal.NumField(); i++ {
			field := dimensionType.Field(i)
			jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]

			if jsonTag == key {
				if dimensionVal.Field(i).CanSet() { //This is compulsary check, check to see if the field is allowed to be set
					// dimensionVal.Field(i).Set(reflect.ValueOf(value).Convert(dimensionVal.Field(i).Type()))
					targetField := dimensionVal.Field(i)
					targetType := targetField.Type()
					//This is the value that we are getting
					valueReflect := reflect.ValueOf(value)

					if targetType.Kind() == reflect.Ptr { //This is the case where our target field is a pointer
						elemType := targetType.Elem() // We need the value/object not the pointer hence we chain .Elem()

						//Check if the value we get can be converted to the value we have in our model (target)
						if !valueReflect.Type().ConvertibleTo(elemType) {
							fmt.Println("Can't Convert the value tp different ")
							return
						}
						//Convert the value
						convertedValue := valueReflect.Convert(elemType)

						//Create a new, pointer to a new value of the element type because we are storing pointers in our model
						newValuePtr := reflect.New(elemType)
						newValuePtr.Elem().Set(convertedValue)
						targetField.Set(newValuePtr)
					} else {
						if !valueReflect.Type().ConvertibleTo(targetType) {
							fmt.Println("Can't convert the values to different type")
							return
						}
						//Set the target field with the value we have as reflect value by converting it to the type of target
						targetField.Set(valueReflect.Convert(targetType))
					}
					break //Move to the next field
				} else {
					fmt.Println("Field Cannot be Set")
					continue
				}
			}
		}
	}

	update_query := `UPDATE dimensions SET length=$1, shoulder=$2, upper_chest=$3, chest=$4, waist=$5, hip=$6, sleeves=$7, neck_front=$8, neck_back=$9, armhole=$10, bottom=$11 WHERE 1=1`

	var update_args []interface{}
	update_args = append(update_args, &existingDimension.Length, &existingDimension.Shoulder, &existingDimension.UpperChest, &existingDimension.Chest, &existingDimension.Waist, &existingDimension.Hip, &existingDimension.Sleeves, &existingDimension.NeckFront, &existingDimension.NeckBack, &existingDimension.Armhole, &existingDimension.Bottom)
	update_args = append(update_args, id)
	// update_query,update_args = addFilter(update_query,name,update_args)
	update_query += ` AND customer_id = $12`
	if name != "" {
		update_args = append(update_args, name)
		query += ` AND customer_name=$13`
	}
	// fmt.Println(len(args))
	_, err = sqlconnect.Dbpool.Exec(context.Background(), update_query, update_args...)
	if err!=nil {
		http.Error(w, fmt.Sprintf("%s",err),http.StatusInternalServerError)
		return
	}
	
	response := struct {
		Status string           `json:"status"`
		Data   models.Dimension `json:"data"`
	}{
		Status: "success",
		Data:   existingDimension,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func addFilter(query string, name string, args []interface{}) (string, []interface{}) {
	query += ` AND customer_id = $1`
	if name != "" {
		args = append(args, name)
		query += ` AND customer_name=$2`
	}
	return query, args
}
