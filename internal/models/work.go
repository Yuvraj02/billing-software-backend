package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkModel struct {
	WorkId        uuid.UUID    `json:"work_id" db:"work_id"`
	CustomerId    int       `json:"customer_id" db:"customer_id"`     //This is going to be Foreign Key Referencing Customers Table
	CustomerName  string    `json:"customer_name" db:"customer_name"` //Needed because it is possible that there are multiple customers on same phone number
	CustomerEmail *string   `json:"customer_email,omitempty" db:"customer_email,omitempty"`
	CustomerPhone string    `json:"customer_phone" db:"customer_phone"`
	WorkStatus    string    `json:"work_status" db:"work_status"`
	Date          time.Time `json:"date,omitempty" db:"date,omitempty"`
	Length        *float32  `json:"length,omitempty" db:"length,omitempty"`
	Shoulder      *float32  `json:"shoulder,omitempty" db:"shoulder,omitempty"`
	UpperChest    *float32  `json:"upper_chest,omitempty" db:"upper_chest,omitempty"`
	Chest         *float32  `json:"chest,omitempty" db:"chest,omitempty"`
	Waist         *float32  `json:"waist,omitempty" db:"waist,omitempty"`
	Hip           *float32  `json:"hip,omitempty" db:"hip,omitempty"`
	Sleeves       *float32  `json:"sleeves,omitempty" db:"sleeves,omitempty"`
	NeckFront     *float32  `json:"neck_front,omitempty" db:"neck_front,omitempty"`
	NeckBack      *float32  `json:"neck_back,omitempty" db:"neck_back,omitempty"`
	Armhole       *float32  `json:"armhole,omitempty" db:"armhole,omitempty"`
	Bottom        *float32  `json:"bottom,omitempty" db:"bottom,omitempty"`
	Category      string    `json:"category,omitempty" db:"category,omitempty"`
}
