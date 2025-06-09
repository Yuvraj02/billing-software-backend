//TODO add filter to request data in ascending order

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
	"time"
	"github.com/jackc/pgx/v5"
)

func GetPendingWork(w http.ResponseWriter, r *http.Request) {

	//get all db columns first using reflect package
	var workModel models.WorkModel
	//Now we create a reflect value of our work model struct
	workModelRefVal := reflect.ValueOf(&workModel).Elem()
	// fmt.Println(workValue)
	workModelType := workModelRefVal.Type()

	var db_columns []string

	query := `SELECT `

	for i := 0; i < workModelRefVal.NumField(); i++ {
		// fmt.Println(workModelType.Field(i).Tag.Get("db"))

		//Extract tag
		db_tag := workModelType.Field(i).Tag.Get("db")
		db_tag = strings.Split(db_tag, ",")[0]
		//Append it to the db_column list
		db_columns = append(db_columns, db_tag)

		query += db_tag
		if i != workModelType.NumField()-1 {
			query += ", "
		}
	}

	query += " FROM worklist WHERE work_status = 'Pending'"

	rows, err := sqlconnect.Dbpool.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Worklist Query Problem", http.StatusInternalServerError)
		return
	}

	workData, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.WorkModel])
	if err != nil {
		http.Error(w, "Error in Collecting Work Rows", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		Count  int    `json:"count"`
		Data   []*models.WorkModel `json:"data"`
	}{
		Status: "success",
		Count:  len(workData),
		Data:   workData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPendingWorkByPhone(w http.ResponseWriter, r *http.Request){

	phoneStr := r.PathValue("phone")
	/*TODO : 

		--------------Implementation left

	*/

}

func AddWork(w http.ResponseWriter, r *http.Request) {

	var newWorkModel models.WorkModel
	json.NewDecoder(r.Body).Decode(&newWorkModel)
	defer r.Body.Close()

	//Add functionality to add new work
	query := `INSERT INTO worklist `
	newWorkModelRefValue := reflect.ValueOf(&newWorkModel).Elem()
	newWorkModelRefType := newWorkModelRefValue.Type()
	//This to be changed later
	var values string = "("
	var db_cols string = "("
	var queryArgs []interface{}

	inc := 1
	for i := 0; i < newWorkModelRefType.NumField(); i++ {

		//This contains value of the fields as we have already dereferenced newWorkModel before using Elem
		targetField := newWorkModelRefValue.Field(i)
		// targetFieldType := targetField.Type()
		targetFieldTag := newWorkModelRefType.Field(i).Tag.Get("db")
		targetFieldTag = strings.Split(targetFieldTag, ",")[0]

		//Even after deref, if there's a field which is pointer, then check again whether the pointer is valid or not
		//If the pointer is not valid, then it means that it is pointing to null values so don't add them
		if targetField.Kind() == reflect.Ptr {

			targetFieldVal := targetField.Elem()

			//if targetFieldVal is again a pointer
			if targetFieldVal.IsValid() { //Check first whether it is valid or not
				db_cols += targetFieldTag + ", "
				values += "$" + strconv.Itoa(inc) + ", "
				inc++
				if targetFieldVal.Kind() == reflect.Float32 {

					queryArgs = append(queryArgs, targetFieldVal.Float())
				} else {
					queryArgs = append(queryArgs, targetFieldVal.String())
				}
			}
		} else {
			db_cols += targetFieldTag + ", "
			values += "$" + strconv.Itoa(inc) + ", "
			inc++
			if targetField.Kind() == reflect.Int {
				queryArgs = append(queryArgs, targetField.Int())
			} else if targetField.Kind() == reflect.String {
				queryArgs = append(queryArgs, targetField.String())
			} else {
				queryArgs = append(queryArgs, targetField.Interface().(time.Time))
			}
		}
	}

	values = values[:len(values)-2]
	values += ")"
	db_cols = db_cols[:len(db_cols)-2]
	db_cols += ")"

	query += db_cols + ` VALUES ` + values
	fmt.Println(queryArgs...)

	// for i := 0; i < len(queryArgs); i++ {
	// 	fmt.Println("Count of ",i+1, "Value : ",queryArgs[i])
	// }

	_, err := sqlconnect.Dbpool.Exec(context.Background(), query, queryArgs...)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}
}
