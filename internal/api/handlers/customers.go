package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/billing-backend/internal/models"
	"restapi/billing-backend/internal/repository/sqlconnect"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	//We want to cancel the context after 2 minutes so that if user gets impatient, then all resources that are held up by this process/routine is cleaned up automatically
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	
	defer cancel()

	rows, _ := sqlconnect.Dbpool.Query(ctx, `SELECT id,name,email,phone,userid,address FROM CUSTOMERS`)

	//Using Collect Rows because it is safer and faster then simple Query and defer rows.Close() automatically 
	//RowToAddOfStructByPos will return pointer to the struct where the values are inserted by position according to our database and our model
	customersList, err := pgx.CollectRows(rows,pgx.RowToAddrOfStructByPos[models.Customer])

	if err!=nil{
		http.Error(w, fmt.Sprintf("%s",err), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	
	response := struct {
		Status string            `json:"status"`
		Count  int               `json:"customers_count"`
		Data   []*models.Customer `json:"data"`
	}{
		Status: "success",
		Count:  len(customersList),
		Data:   customersList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}


func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	
	idStr := r.PathValue("id")
	id,err := strconv.Atoi(idStr)

	if err!=nil{
		fmt.Printf("%s",err)
		return
	}

	query := `SELECT id,name,email,phone,userid,address FROM customers WHERE id = $1`
	row := sqlconnect.Dbpool.QueryRow(context.Background(), query,id)

	var customer models.Customer
	err = row.Scan(&customer.Id,&customer.Name,&customer.Email,&customer.Phone, &customer.UserID, &customer.Address)
	if err!=nil{
		http.Error(w,fmt.Sprintf("%s",err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		Data models.Customer `json:"data"`
	}{
		Status: "Success",
		Data: customer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func AddCustomer(w http.ResponseWriter, r *http.Request) {

	var newCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)

	defer r.Body.Close()

	query := `INSERT INTO customers (id,name,email,phone,userid,address) VALUES (DEFAULT, $1,$2,$3,$4,$5)`
	_,err := sqlconnect.Dbpool.Exec(context.Background(), query,newCustomer.Name, newCustomer.Email, newCustomer.Phone, newCustomer.UserID, newCustomer.Address)

	if err!=nil{
		http.Error(w, "Error in inserting cutomer to database", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		Data models.Customer `json:"data_added"`
	} {
		Status: "success",
		Data: newCustomer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}