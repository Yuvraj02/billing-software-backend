package models

/*

	When API will be called on the frontent
	This will return all non null values
	These Non-Null values will be the name of the Dimension so no additional mapping will be needed in frontend

*/

type Category struct {
	CategoryId   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Shoulder     *string `json:"shoulder,omitempty"`
	UpperChest   *string `json:"upper_chest,omitempty"`
	Chest        *string `json:"chest,omitempty"`
	Waist        *string `json:"waist,omitempty"`
	Hip          *string `json:"hip,omitempty"`
	Sleeves      *string `json:"sleeves,omitempty"`
	NeckFront    *string `json:"neck_front,omitempty"`
	NeckBack     *string `json:"neck_back,omitempty"`
	Armhole      *string `json:"armhole,omitempty"`
	Length       *string `json:"length,omitempty"`
	Bottom       *string `json:"bottom,omitempty"`
}
