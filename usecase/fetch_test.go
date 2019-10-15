package usecase

import (
	"projectapi/models"
	postgresrepo "projectapi/repository/postgres"
	restapirepo "projectapi/repository/restapi"
	"reflect"
	"testing"
)

func TestGarageServiceStruct_GetGarageDetail(t *testing.T) {
	type fields struct {
		Postgres postgresrepo.RepoPostgres
		Cardata  restapirepo.RepoCAR
	}
	type args struct {
		IDGarage int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantRetData Garage
	}{
		{
			name: "Testcase 1",
			fields: fields{
				Postgres: &postgresrepo.RepoPostgresMock{
					GetGaragesByIDFunc: func(garageID int) models.Garages {
						return models.Garages{
							IDGarage:   2,
							GarageNm:   "E5 Garage",
							Latitude:   "0",
							Longtitude: "0",
						}
					},
					GetPositionByGarageIDFunc: func(garageID int) []models.Position {
						return []models.Position{
							models.Position{
								IDPosition:   4,
								IDCar:        1,
								IDGarage:     2,
								PositionName: "Slot A",
							},
							models.Position{
								IDPosition:   5,
								IDCar:        2,
								IDGarage:     2,
								PositionName: "Slot B",
							},
						}
					},
					GetCarPositionFunc: func(idcar string) []models.CarPosition {
						return []models.CarPosition{
							{
								GarageName:   "E4 Garage",
								Longtitude:   "1.232",
								Latitude:     "0.12",
								PositionName: "Slot 1",
							},
						}
					},
				},
				Cardata: &restapirepo.RepoCARMock{
					GetGarageStatusFunc: func(IDgarage int) map[int]models.GarageStatus {
						return map[int]models.GarageStatus{
							1: models.GarageStatus{
								OwnerName: "Ceki",
								CarName:   "J",
								IDCar:     "1",
							},
							2: models.GarageStatus{
								OwnerName: "Ceko",
								CarName:   "Kijang",
								IDCar:     "2",
							},
						}
					},
				},
			},
			args: args{
				IDGarage: 2,
			},
			wantRetData: Garage{
				IDGarage:    2,
				GarageName:  "E5 Garage",
				GarageOwner: "",
				Position: []positionDetail{
					positionDetail{
						IDPosition: 4,
						SlotName:   "Slot A",
						CarName:    "J",
					},
					positionDetail{
						IDPosition: 5,
						SlotName:   "Slot B",
						CarName:    "Kijang",
					},
				},
			},
		},
		{
			name: "Testcase 2 Kosong",
			fields: fields{
				Postgres: &postgresrepo.RepoPostgresMock{
					GetGaragesByIDFunc: func(garageID int) models.Garages {
						return models.Garages{}
					},
					GetPositionByGarageIDFunc: func(garageID int) []models.Position {
						return []models.Position{}
					},
				},
				Cardata: &restapirepo.RepoCARMock{
					GetGarageStatusFunc: func(IDgarage int) map[int]models.GarageStatus {
						return map[int]models.GarageStatus{}
					},
				},
			},
			args: args{
				IDGarage: 0,
			},
			wantRetData: Garage{},
		},
		{
			name: "Testcase 3 nil",
			fields: fields{
				Postgres: &postgresrepo.RepoPostgresMock{
					GetGaragesByIDFunc: func(garageID int) models.Garages {
						return models.Garages{}
					},
					GetPositionByGarageIDFunc: func(garageID int) []models.Position {
						return nil
					},
				},
				Cardata: &restapirepo.RepoCARMock{
					GetGarageStatusFunc: func(IDgarage int) map[int]models.GarageStatus {
						return nil
					},
				},
			},
			args: args{
				IDGarage: 0,
			},
			wantRetData: Garage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := &GarageServiceStruct{
				Postgres: tt.fields.Postgres,
				Cardata:  tt.fields.Cardata,
			}
			if gotRetData := G.GetGarageDetail(tt.args.IDGarage); !reflect.DeepEqual(gotRetData, tt.wantRetData) {
				t.Errorf("GarageServiceStruct.GetGarageDetail() = %v, want %v", gotRetData, tt.wantRetData)
			}
		})
	}
}

func TestGarageServiceStruct_GetCarPosition(t *testing.T) {
	type fields struct {
		Postgres postgresrepo.RepoPostgres
		Cardata  restapirepo.RepoCAR
	}
	type args struct {
		idcar string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.CarPosition
	}{
		{
			name: "Testcase 2 Kosong",
			fields: fields{
				Postgres: &postgresrepo.RepoPostgresMock{
					GetCarPositionFunc: func(idcar string) []models.CarPosition {
						return []models.CarPosition{}
					},
				},
			},
			args: args{
				idcar: "0",
			},
			want: []models.CarPosition{},
		},
		{
			name: "Testcase 1",
			fields: fields{
				Postgres: &postgresrepo.RepoPostgresMock{
					GetCarPositionFunc: func(idcar string) []models.CarPosition {
						return []models.CarPosition{
							{
								GarageName:   "E4 Garage",
								Longtitude:   "1.232",
								Latitude:     "0.12",
								PositionName: "Slot 1",
							},
						}
					},
				},
			},
			args: args{
				idcar: "2",
			},
			want: []models.CarPosition{
				{
					GarageName:   "E4 Garage",
					Longtitude:   "1.232",
					Latitude:     "0.12",
					PositionName: "Slot 1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := &GarageServiceStruct{
				Postgres: tt.fields.Postgres,
				Cardata:  tt.fields.Cardata,
			}
			if got := G.GetCarPosition(tt.args.idcar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GarageServiceStruct.GetCarPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
