package models

/*

	This model will contain the dimensions for a customer
	This will be helpful when we store several customers in a database
	Fields are set as pointers because they can be null
*/

type Dimension struct {
	CustomerId    int      `json:"customer_id"` //This is going to be Foreign Key Referencing Customers Table
	CustomerName  string 	`json:"customer_name"` //Needed because it is possible that there are multiple customers on same phone number
	CustomerPhone string   `json:"customer_phone"`
	Length        *float32 `json:"length,omitempty"`
	Shoulder      *float32 `json:"shoulder,omitempty"`
	UpperChest    *float32 `json:"upper_chest,omitempty"`
	Chest         *float32 `json:"chest,omitempty"`
	Waist         *float32 `json:"waist,omitempty"`
	Hip           *float32 `json:"hip,omitempty"`
	Sleeves       *float32 `json:"sleeves,omitempty"`
	NeckFront     *float32 `json:"neck_front,omitempty"`
	NeckBack      *float32 `json:"neck_back,omitempty"`
	Armhole       *float32 `json:"armhole,omitempty"`
	Bottom        *float32 `json:"bottom,omitempty"`
}
