package usecase

import (
	"projectapi/models"
)

type positionDetail struct {
	IDPosition int
	SlotName   string
	CarName    string
}

// Garage is Struct who show garage and position
type Garage struct {
	IDGarage    int
	GarageName  string
	GarageOwner string
	Position    []positionDetail
}

// GetGarageDetail get garage Detil
func (G *GarageServiceStruct) GetGarageDetail(IDGarage int) (retData Garage) {

	garageData := G.Postgres.GetGaragesByID(IDGarage)
	garageDetail := G.Cardata.GetGarageStatus(IDGarage)
	garageSlot := G.Postgres.GetPositionByGarageID(IDGarage)

	retData.IDGarage = IDGarage
	retData.GarageName = garageData.GarageNm
	retData.GarageOwner = garageDetail[0].OwnerName

	for _, item := range garageSlot {
		var each positionDetail
		each.IDPosition = item.IDPosition
		each.SlotName = item.PositionName
		each.CarName = garageDetail[item.IDCar].CarName
		retData.Position = append(retData.Position, each)
	}
	return retData
}

// GetCarPosition get position of cars
func (G *GarageServiceStruct) GetCarPosition(idcar string) []models.CarPosition {
	var positionDatas []models.CarPosition
	positionDatas = G.Postgres.GetCarPosition(idcar)
	return positionDatas
}
