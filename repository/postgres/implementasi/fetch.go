package postgres

import (
	"fmt"
	"projectapi/models"
)

// GetGaragesByID is to get garages by ID
func (repo *Repo) GetGaragesByID(garageID int) models.Garages {
	var garageData models.Garages
	rows := repo.prepGetGarageByID.QueryRow(garageID)
	rows.Scan(&garageData.IDGarage, &garageData.GarageNm, &garageData.Latitude, &garageData.Longtitude)
	return garageData
}

// GetPositionByGarageID is to get position by garage id
func (repo *Repo) GetPositionByGarageID(garageID int) []models.Position {
	var positionData []models.Position
	rows, err := repo.prepGetPositionByGarageID.Query(garageID)
	if err != nil {
		fmt.Println("Query failed ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var each = models.Position{}
		rows.Scan(&each.IDPosition, &each.IDGarage, &each.IDCar, &each.PositionName)
		positionData = append(positionData, each)
	}
	return positionData
}

// GetCarPosition get position of cars
func (repo *Repo) GetCarPosition(idcar string) []models.CarPosition {
	var positionDatas []models.CarPosition
	rows, err := repo.prepGetCarPosition.Query(idcar)
	if err != nil {
		fmt.Println("Query failed ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var each models.CarPosition
		rows.Scan(&each.GarageName, &each.Latitude, &each.Longtitude, &each.PositionName)
		positionDatas = append(positionDatas, each)
	}
	return positionDatas
}