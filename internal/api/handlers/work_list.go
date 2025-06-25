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

	"github.com/google/uuid"
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

	query += " FROM worklist WHERE work_status = 'Pending' ORDER BY date DESC"

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

func GetPendingWorkByName(w http.ResponseWriter, r *http.Request){

	nameStr := r.PathValue("name")

	//`SELECT * FROM worklist WHERE customer_name = nameStr`

		query := `SELECT `
		db_cols := ""
		
		var workModel models.WorkModel

		workModelRefVal := reflect.ValueOf(&workModel).Elem()
		workModelRefType := workModelRefVal.Type()

		for i:=0; i<workModelRefType.NumField(); i++ {

			dbTag := workModelRefType.Field(i).Tag.Get("db")
			dbTag = strings.Split(dbTag,",")[0]
			// fmt.Println(dbTag)
			db_cols += dbTag

			if(i!=workModelRefType.NumField()-1){
				db_cols += ", "
			}
		}

		// fmt.Println(db_cols)
		query += db_cols + ` FROM worklist WHERE customer_name = $1 AND work_status = 'Pending'`
		
		rows, err := sqlconnect.Dbpool.Query(context.Background(), query, nameStr)
		if err!=nil{
			fmt.Println(fmt.Sprintf("%s", err))
			return
		}
		
		workList, err := pgx.CollectRows(rows,pgx.RowToAddrOfStructByPos[models.WorkModel])
		if err!=nil{
			fmt.Println(err)
			return
		}

		response := struct {
			Status string `json:"status"`
			Data []*models.WorkModel `json:"data"`
		}{
			Status: "success",
			Data: workList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
}

func GetPendingWorkByPhone(w http.ResponseWriter, r *http.Request){

	phoneStr := r.PathValue("phone")
	/*TODO : 

		--------------Implementation left
	*/
	// query := `SELECT work_id, customer_id, customer_name, customer_phone, work_status, date FROM worklist WHERE customer_phone = $1 AND work_status = 'Pending'`
	// row := sqlconnect.Dbpool.QueryRow(context.Background(), query, phoneStr)
	// var workModel models.WorkModel
	// err := row.Scan(&workModel.WorkId,&workModel.CustomerId, &workModel.CustomerName, &workModel.CustomerPhone, &workModel.WorkStatus, &workModel.Date)
	// if err != nil {
	// 	fmt.Println("Error in Finding Pending Work for customer")
	// 	return
	// }

	// response := struct {
	// 	Status string `json:"status"`
	// 	Data models.WorkModel `json:"data"`
	// }{
	// 	Status: "success",
	// 	Data: workModel,
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)

	query:=`SELECT `

	db_cols := ""

	var workModel models.WorkModel
	workModelRefVal := reflect.ValueOf(&workModel).Elem()
	workModelRefType := workModelRefVal.Type()

	for i:=0 ; i<workModelRefType.NumField(); i++ {

		targetField := workModelRefType.Field(i)
		targetFieldTag := targetField.Tag.Get("db")

		targetFieldTag = strings.Split(targetFieldTag, ",")[0]

		db_cols += targetFieldTag

		if i!=workModelRefType.NumField()-1 {
			db_cols += ", "
		}
		
	}

	query += db_cols + ` FROM worklist WHERE customer_phone = $1 AND work_status = 'Pending'`	

	rows ,err := sqlconnect.Dbpool.Query(context.Background(), query, phoneStr)
	if err!=nil{
		fmt.Println(query)
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}

	userList , err:= pgx.CollectRows(rows, pgx.RowToStructByPos[models.WorkModel])
	if err!=nil{
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}

	response := struct {
		Status string `json:"sucess"`
		Data []models.WorkModel `json:"data"`
	}{
		Status: "success",
		Data: userList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	
}

func GetPendingWorkById(w http.ResponseWriter, r *http.Request){

	id := r.PathValue("work_id")

	query := `SELECT `
	//Build Query
	var workModel models.WorkModel
	workModelRefVal := reflect.ValueOf(&workModel).Elem()
	workModelRefType := workModelRefVal.Type()
	db_cols := ""
	for i :=0 ; i<workModelRefType.NumField(); i++{

		// targetField := workModelRefVal.Field(i)
		targetFieldTag := workModelRefType.Field(i).Tag.Get("db")
		targetFieldTag = strings.Split(targetFieldTag, ",")[0]
		
		db_cols += targetFieldTag 
		if(i!=workModelRefType.NumField()-1){
			db_cols += ", "
		}
	}

	query+=db_cols + ` FROM worklist WHERE work_id = $1`

	row := sqlconnect.Dbpool.QueryRow(context.Background(), query, id)

	//Get Address of fields in a list
	var workModelArgs []interface{}
	for i := 0; i < workModelRefType.NumField(); i++ {
		//Get Address of the fields
		workModelArgs = append(workModelArgs, workModelRefVal.Field(i).Addr().Interface())
	}

	//Scan All the values in those addresses
	err := row.Scan(workModelArgs...)
	if err!=nil {
		http.Error(w,fmt.Sprintf("%s", err),http.StatusInternalServerError)
		return
	}

	response := struct{
		Status string `json:"status"`
		Data models.WorkModel	`json:"data"`
	}{
		Status: "success",
		Data: workModel,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
			} else if targetField.Kind() == reflect.Array{
				queryArgs = append(queryArgs, targetField.Interface().(uuid.UUID))
			}else{
				queryArgs = append(queryArgs, targetField.Interface().(time.Time))
			}
		}
	}

	values = values[:len(values)-2]
	values += ")"
	db_cols = db_cols[:len(db_cols)-2]
	db_cols += ")"

	query += db_cols + ` VALUES ` + values
	// fmt.Println(queryArgs...)

	_, err := sqlconnect.Dbpool.Exec(context.Background(), query, queryArgs...)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}
}

func PatchWork(w http.ResponseWriter, r *http.Request){

	idStr := r.PathValue("id")

	query := `UPDATE worklist SET work_status = 'Completed' WHERE work_id = $1`

	_, err := sqlconnect.Dbpool.Exec(context.Background(), query, idStr)

	if err!=nil{
		http.Error(w,fmt.Sprintf("%s",err), http.StatusInternalServerError)
		return
	}
}

func GetCompletedWork(w http.ResponseWriter, r *http.Request){

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

	query += " FROM worklist WHERE work_status = 'Completed' ORDER BY date DESC"

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

func GetCompletedWorkByPhone(w http.ResponseWriter, r *http.Request){

	phoneStr := r.PathValue("phone")

	query:=`SELECT `

	db_cols := ""

	var workModel models.WorkModel
	workModelRefVal := reflect.ValueOf(&workModel).Elem()
	workModelRefType := workModelRefVal.Type()

	for i:=0 ; i<workModelRefType.NumField(); i++ {

		targetField := workModelRefType.Field(i)
		targetFieldTag := targetField.Tag.Get("db")

		targetFieldTag = strings.Split(targetFieldTag, ",")[0]

		db_cols += targetFieldTag

		if i!=workModelRefType.NumField()-1 {
			db_cols += ", "
		}
		
	}

	query += db_cols + ` FROM worklist WHERE customer_phone = $1 AND work_status = 'Completed'`	

	rows ,err := sqlconnect.Dbpool.Query(context.Background(), query, phoneStr)
	if err!=nil{
		fmt.Println(query)
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}

	userList , err:= pgx.CollectRows(rows, pgx.RowToStructByPos[models.WorkModel])
	if err!=nil{
		fmt.Println(fmt.Sprintf("%s", err))
		return
	}

	response := struct {
		Status string `json:"sucess"`
		Data []models.WorkModel `json:"data"`
	}{
		Status: "success",
		Data: userList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetCompletedWorkByName(w http.ResponseWriter, r *http.Request){

	nameStr := r.PathValue("name")

	//`SELECT * FROM worklist WHERE customer_name = nameStr`

		query := `SELECT `
		db_cols := ""
		
		var workModel models.WorkModel

		workModelRefVal := reflect.ValueOf(&workModel).Elem()
		workModelRefType := workModelRefVal.Type()

		for i:=0; i<workModelRefType.NumField(); i++ {

			dbTag := workModelRefType.Field(i).Tag.Get("db")
			dbTag = strings.Split(dbTag,",")[0]
			// fmt.Println(dbTag)
			db_cols += dbTag

			if(i!=workModelRefType.NumField()-1){
				db_cols += ", "
			}
		}

		query += db_cols + ` FROM worklist WHERE customer_name = $1 AND work_status = 'Completed'`
		
		rows, err := sqlconnect.Dbpool.Query(context.Background(), query, nameStr)
		if err!=nil{
			fmt.Println(fmt.Sprintf("%s", err))
			return
		}
		
		workList, err := pgx.CollectRows(rows,pgx.RowToAddrOfStructByPos[models.WorkModel])
		if err!=nil{
			fmt.Println(err)
			return
		}

		response := struct {
			Status string `json:"status"`
			Data []*models.WorkModel `json:"data"`
		}{
			Status: "success",
			Data: workList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
}

