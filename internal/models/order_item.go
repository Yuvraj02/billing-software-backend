package models

type OrderItem struct {
	ProductID    int `json:"product_id"`
	CustomerID   int `json:"customer_id"`
	OrderedUnits int `json:"ordered_units"`
	OrderPrce    int `json:"order_price"`
}
