package models

type Dimension struct {
	CustomerId int `json:"customer_id"`
	Shoulder   int `json:"shoulder"`
	UpperChest int `json:"upper_chest"`
	Chest      int `json:"chest"`
	Waist      int `json:"waist"`
	Hip        int `json:"hip"`
	Sleeves    int `json:"sleeves"`
	NeckFront  int `json:"neck_front"`
	NeckBack   int `json:"neck_back"`
	Armhole    int `json:"armhole"`
}
