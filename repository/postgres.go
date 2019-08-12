package db

import (
	"database/sql"
	"fmt"
	md "projectapi/models"

	_ "github.com/lib/pq"
)

// Conn Open connection to server
func Conn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:changeme@localhost/garages?sslmode=disable")
	if err != nil {
		fmt.Println(`Could not connect to db`)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(`Could not connect to db`)
		return nil, err
	}

	return db, nil
}

// GetGarages is to get garages list all
func GetGarages(db *sql.DB) []md.Garages {
	var garageDatas []md.Garages
	sqlStatement := "select id_garage, garage_nm, latitude, longtitude from garages"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Query faiiled ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var each = md.Garages{}
		rows.Scan(&each.IDGarage, &each.GarageNm, &each.Latitude, &each.Longtitude)
		garageDatas = append(garageDatas, each)
	}
	return garageDatas
}

// GetGaragesByID is to get garages by ID
func GetGaragesByID(db *sql.DB, garageID int) md.Garages {
	var garageData md.Garages
	sqlStatement := "select id_garage, garage_nm, latitude, longtitude from garages where id_garage = $1"
	rows := db.QueryRow(sqlStatement, garageID)
	rows.Scan(&garageData.IDGarage, &garageData.GarageNm, &garageData.Latitude, &garageData.Longtitude)
	return garageData
}

// GetPositionByGarageID is to get position by garage id
func GetPositionByGarageID(db *sql.DB, garageID int) []md.Position {
	var positionData []md.Position
	sqlStatement := fmt.Sprintf("select id_position, id_garage, id_car, position_nm from position where id_garage = %d", garageID)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Query failed ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var each = md.Position{}
		rows.Scan(&each.IDPosition, &each.IDGarage, &each.IDCar, &each.PositionName)
		positionData = append(positionData, each)
	}
	return positionData
}

// GetCarPosition get position of cars
func GetCarPosition(db *sql.DB, idcar string) []md.CarPosition {
	var positionDatas []md.CarPosition
	sqlStatement := "select b.garage_nm, b.longtitude, b.latitude, a.position_nm from position a left join garages b on a.id_garage = b.id_garage where a.id_car = $1"
	rows, err := db.Query(sqlStatement, idcar)
	if err != nil {
		fmt.Println("Query failed ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var each md.CarPosition
		rows.Scan(&each.GarageName, &each.Latitude, &each.Longtitude, &each.PositionName)
		positionDatas = append(positionDatas, each)
	}
	return positionDatas
}
