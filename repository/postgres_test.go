package db

import (
	"database/sql"
	md "projectapi/models"
	"reflect"
	"testing"

	asr "github.com/stretchr/testify/assert"
)

func Test_GetGaragesByID(t *testing.T) {
	dbs, _ := Conn()
	defer dbs.Close()
	type args struct {
		db       *sql.DB
		IDGarage int
	}
	tests := []struct {
		name        string
		arguments   args
		wantIsExist md.Garages
	}{
		{
			name: "Test Case 1 garage Exist",
			arguments: args{
				db:       dbs,
				IDGarage: 1,
			},
			wantIsExist: md.Garages{
				IDGarage:   1,
				GarageNm:   "E4 Garage",
				Latitude:   "1.232",
				Longtitude: "0.12",
			},
		},
		{
			name: "Test Case 2 garage Not Exist",
			arguments: args{
				db:       dbs,
				IDGarage: 1000000000000,
			},
			wantIsExist: md.Garages{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if gotExist := GetGaragesByID(testcase.arguments.db, testcase.arguments.IDGarage); !reflect.DeepEqual(gotExist, testcase.wantIsExist) {
				t.Errorf(" Checking function GetGaragesByID failed %v %v", gotExist, testcase.wantIsExist)
			}
		})
	}
}

func Test_GetPositionByGarageID(t *testing.T) {
	dbs, _ := Conn()
	defer dbs.Close()
	type args struct {
		db       *sql.DB
		IDGarage int
	}
	tests := []struct {
		name        string
		arguments   args
		wantIsExist []md.Position
	}{
		{
			name: "Test Case 1 garage and position Exist",
			arguments: args{
				db:       dbs,
				IDGarage: 1,
			},
			wantIsExist: []md.Position{
				md.Position{
					IDPosition:   1,
					IDGarage:     1,
					IDCar:        1,
					PositionName: "Slot 1",
				},
			},
		},
		{
			name: "Test Case 2 garage and position Not Exist",
			arguments: args{
				db:       dbs,
				IDGarage: 1000000000000,
			},
			wantIsExist: []md.Position{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if gotExist := GetPositionByGarageID(testcase.arguments.db, testcase.arguments.IDGarage); reflect.TypeOf(gotExist) != reflect.TypeOf(testcase.wantIsExist) {
				t.Errorf(" Checking function GetGaragesByID failed %v %v", gotExist, testcase.wantIsExist)
			}
		})
	}
}

func Test_GetCarPosition(t *testing.T) {
	dbs, _ := Conn()
	defer dbs.Close()
	type args struct {
		db    *sql.DB
		IDCar string
	}
	tests := []struct {
		name        string
		arguments   args
		wantIsExist []md.CarPosition
	}{
		{
			name: "Test Case 1 car position Exist",
			arguments: args{
				db:    dbs,
				IDCar: "1",
			},
			wantIsExist: []md.CarPosition{
				md.CarPosition{
					GarageName:   "E4 Garage",
					Latitude:     "1.232",
					Longtitude:   "0.12",
					PositionName: "Slot 1",
				},
			},
		},
		{
			name: "Test Case 2 car position Not Exist",
			arguments: args{
				db:    dbs,
				IDCar: "100",
			},
			wantIsExist: []md.CarPosition{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if gotExist := GetCarPosition(testcase.arguments.db, testcase.arguments.IDCar); reflect.TypeOf(gotExist) != reflect.TypeOf(testcase.wantIsExist) {
				t.Errorf(" Checking function GetGaragesByID failed %v %v", gotExist, testcase.wantIsExist)
			}
		})
	}
}

func Test_GetGarages(t *testing.T) {
	dbs, _ := Conn()
	defer dbs.Close()
	assert := asr.New(t)
	type args struct {
		db    *sql.DB
		IDCar string
	}
	tests := []struct {
		name        string
		arguments   args
		wantIsExist md.Garages
	}{
		{
			name: "Test Case 1 car position Exist",
			arguments: args{
				db: dbs,
			},
			wantIsExist: md.Garages{
				IDGarage:   1,
				GarageNm:   "E4 Garage",
				Latitude:   "1.232",
				Longtitude: "0.12",
			},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			gotExist := GetGarages(testcase.arguments.db)
			assert.Contains(gotExist, testcase.wantIsExist)
		})
	}
}

func Test_Conn(t *testing.T) {
	dbs, _ := Conn()
	t.Run("Test Connection", func(t *testing.T) {
		if err := dbs.Ping(); err != nil {
			t.Errorf("Error Creating Connection %v ", err)
		}
	})
}
