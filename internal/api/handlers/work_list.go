package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"restapi/billing-backend/internal/models"
	"restapi/billing-backend/internal/repository/sqlconnect"
	"strings"

	"github.com/jackc/pgx/v5"
)

func GetPendingWork(w http.ResponseWriter, r *http.Request){

	//get all db columns first using reflect package
	var workModel models.WorkModel
	//Now we create a reflect value of our work model struct
	workModelRefVal := reflect.ValueOf(&workModel).Elem()
	// fmt.Println(workValue)
	workModelType := workModelRefVal.Type()

	var db_columns []string 
	
	query := `SELECT `

	for i:=0 ; i<workModelRefVal.NumField();i++{
		// fmt.Println(workModelType.Field(i).Tag.Get("db"))
		
		//Extract tag
		db_tag := workModelType.Field(i).Tag.Get("db")
		db_tag = strings.Split(db_tag,",")[0]
		//Append it to the db_column list
		db_columns = append(db_columns,db_tag)
		
		query += db_tag 
		if(i!=workModelType.NumField()-1){
				query += ", "
		}
	}

	query += " FROM worklist WHERE work_status = 'Pending'"

	rows, err := sqlconnect.Dbpool.Query(context.Background(), query)
	if err!=nil{
		http.Error(w,"Worklist Query Problem",http.StatusInternalServerError)
		return
	}

	workData , err := pgx.CollectRows(rows,pgx.RowToAddrOfStructByPos[models.WorkModel])
	if err != nil{
		http.Error(w,"Error in Collecting Work Rows", http.StatusInternalServerError)
		return 
	}

	response := struct {
		Status string `json:"status"`
		Count int `json:"count"`
		Data []*models.WorkModel
	}{
		Status: "success",
		Count: len(workData),
		Data: workData,
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func AddWork(w http.ResponseWriter, r *http.Request){

	var newWorkModel models.WorkModel
	json.NewDecoder(r.Body).Decode(&newWorkModel)
	defer r.Body.Close()

	//Add functionality to add new work

}
