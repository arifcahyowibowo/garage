package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	md "projectapi/models"
	"strconv"
)

// GetGarageStatus get garage status from other service
func GetGarageStatus(IDgarage int) map[int]md.GarageStatus {
	var garageStatusArray []md.GarageStatus
	garageStatusMap := make(map[int]md.GarageStatus)
	url := fmt.Sprintf("http://172.31.4.92:8980/getgeragestatus?id_Gerage=%d", IDgarage)
	res, err := http.Get(url)
	defer res.Body.Close()
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
