package models

type Customer struct {
	Id      int     `json:"customer_id"`
	Name    string  `json:"customer_name"`
	Email   *string `json:"customer_email"`
	Phone   string  `json:"customer_ph"`
	UserID  int     `json:"user_id"`
	Address *string `json:"customer_address"`
}
