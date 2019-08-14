package delivery

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"projectapi/models"
	postgresrepo "projectapi/repository/postgres"
	restapirepo "projectapi/repository/restapi"
	"projectapi/usecase"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDelivery_GetCarPositionHandler(t *testing.T) {
	type fields struct {
		Service *usecase.GarageServiceStruct
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		wantCode int
		wantData string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				Service: &usecase.GarageServiceStruct{
					&postgresrepo.RepoPostgresMock{
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
					&restapirepo.RepoCARMock{
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
			},
			wantCode: 200,
			wantData: `[{"GarageName":"E4 Garage","Longtitude":"1.232","Latitude":"0.12","PositionName":"Slot 1"}]`,
		},

		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := Delivery{
				Service: tt.fields.Service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/api/getcarposition/{idcar}", D.GetCarPositionHandler).Methods("GET")
			recorder := httptest.NewRecorder()
			RequestGarage, _ := http.NewRequest("GET", "/api/getcarposition/2", nil)
			router.ServeHTTP(recorder, RequestGarage)
			resData, _ := ioutil.ReadAll(recorder.Body)
			assert.Equal(t, tt.wantCode, recorder.Code)
			assert.Equal(t, tt.wantData, string(resData))
		})
	}
}

func TestDelivery_GetGarageDetailHandler(t *testing.T) {
	type fields struct {
		Service *usecase.GarageServiceStruct
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantData string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				Service: &usecase.GarageServiceStruct{
					&postgresrepo.RepoPostgresMock{
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
					&restapirepo.RepoCARMock{
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
			},
			wantCode: 200,
			wantData: `{"IDGarage":2,"GarageName":"E5 Garage","GarageOwner":"","Position":[{"IDPosition":4,"SlotName":"Slot A","CarName":"J"},{"IDPosition":5,"SlotName":"Slot B","CarName":"Kijang"}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := Delivery{
				Service: tt.fields.Service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/api/getgaragedetail/{idgarage}", D.GetGarageDetailHandler).Methods("GET")
			recorder := httptest.NewRecorder()
			RequestGarage, _ := http.NewRequest("GET", "/api/getgaragedetail/2", nil)
			router.ServeHTTP(recorder, RequestGarage)
			resData, _ := ioutil.ReadAll(recorder.Body)
			assert.Equal(t, tt.wantCode, recorder.Code)
			assert.Equal(t, tt.wantData, string(resData))
		})
	}
}

func TestDelivery_CreateRouter(t *testing.T) {
	type fields struct {
		Service *usecase.GarageServiceStruct
	}
	tests := []struct {
		name   string
		fields fields
		want   *mux.Router
	}{
		{
			name: "Testcase 1",
			fields: fields{
				Service: &usecase.GarageServiceStruct{
					&postgresrepo.RepoPostgresMock{
						GetGaragesByIDFunc: func(garageID int) models.Garages {
							return models.Garages{}
						},
						GetPositionByGarageIDFunc: func(garageID int) []models.Position {
							return []models.Position{}
						},
						GetCarPositionFunc: func(idcar string) []models.CarPosition {
							return []models.CarPosition{}
						},
					},
					&restapirepo.RepoCARMock{
						GetGarageStatusFunc: func(IDgarage int) map[int]models.GarageStatus {
							return map[int]models.GarageStatus{}
						},
					},
				},
			},
			want: &mux.Router{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			D := Delivery{
				Service: tt.fields.Service,
			}
			if got := D.CreateRouter(); !assert.IsType(t, got, tt.want) {
				t.Errorf("Delivery.CreateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
