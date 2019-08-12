package models

// Garages models for table garages
type Garages struct {
	IDGarage   int    `json:"IDGarage"`
	GarageNm   string `json:"GarageNm"`
	Latitude   string `json:"Latitude"`
	Longtitude string `json:"Longtitude"`
}

// Position models for table position
type Position struct {
	IDPosition   int    `json:"IDPosition"`
	IDGarage     int    `json:"IDGarage"`
	IDCar        int    `json:"IDCar"`
	PositionName string `json:"PositionName"`
}

// CarPosition models for result in getGarageDetail
type CarPosition struct {
	GarageName   string `json:"GarageName"`
	Longtitude   string `json:"Longtitude"`
	Latitude     string `json:"Latitude"`
	PositionName string `json:"PositionName"`
}

// GarageStatus is Struct to parse data from API
type GarageStatus struct {
	OwnerName string `json:"ownerName"`
	CarName   string `json:"CarName"`
	IDCar     string `json:"IdCar"`
}
