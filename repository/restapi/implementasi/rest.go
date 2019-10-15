package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projectapi/models"
	"strconv"
)

// GetGarageStatus get garage status from other service
func (R *RestAPI) GetGarageStatus(IDgarage int) map[int]models.GarageStatus {
	var garageStatusArray []models.GarageStatus
	garageStatusMap := make(map[int]models.GarageStatus)
	url := fmt.Sprintf("%s?%s=%d", R.baseURL, R.endPoint, IDgarage)
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	json.NewDecoder(res.Body).Decode(&garageStatusArray)
	for _, item := range garageStatusArray {
		carID, _ := strconv.Atoi(item.IDCar)
		fmt.Println(carID)
		garageStatusMap[carID] = item
	}
	fmt.Println(garageStatusArray)
	return garageStatusMap
}
