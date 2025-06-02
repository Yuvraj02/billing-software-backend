package models

type Order struct {
	OrderId    int `json:"order_id"`
	CustomerId int `json:"customer_id"`
	TotalPrice      int `json:"order_price"`
	
}
