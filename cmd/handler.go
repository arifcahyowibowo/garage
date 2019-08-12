package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "projectapi/repository"
	"strconv"

	"github.com/gorilla/mux"
)

// GaragesHandler handler for reques garages list
func GaragesHandler(w http.ResponseWriter, r *http.Request) {
	dbs, _ := db.Conn()
	defer dbs.Close()
	garageDatas := db.GetGarages(dbs)
	jsonGarages, err := json.Marshal(garageDatas)
	if err != nil {
		log.Fatal("failed to fetch json")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonGarages))
}

// GetCarPositionHandler handler for reques garages list
func GetCarPositionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idcar := params["idcar"]
	dbs, _ := db.Conn()
	defer dbs.Close()
	carPositionDatas := db.GetCarPosition(dbs, idcar)
	jsonCarPositionDatas, err := json.Marshal(carPositionDatas)
	if err != nil {
		log.Fatal("failed to fetch json")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonCarPositionDatas))
}

// GetGarageDetail handler for reques garages detail
func GetGarageDetail(w http.ResponseWriter, r *http.Request) {
	type positionDetail struct {
		IDPosition int
		SlotName   string
		CarName    string
	}
	type garage struct {
		IDGarage    int
		GarageName  string
		GarageOwner string
		Position    []positionDetail
	}
	var retData garage
	params := mux.Vars(r)
	dbs, err := db.Conn()
	if err != nil {
		fmt.Println(err)
	}
	defer dbs.Close()
	IDGarage, _ := strconv.Atoi(params["idgarage"])
	garageData := db.GetGaragesByID(dbs, IDGarage)
	garageDetail := db.GetGarageStatus(IDGarage)
	garageSlot := db.GetPositionByGarageID(dbs, IDGarage)

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
	jsonRetData, _ := json.Marshal(retData)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonRetData))
}
