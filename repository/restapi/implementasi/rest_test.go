package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projectapi/models"
	"reflect"
	"testing"
)

func TestRestAPI_GetGarageStatus(t *testing.T) {
	type fields struct {
		baseURL  string
		endPoint string
	}
	type args struct {
		IDgarage int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]models.GarageStatus
	}{
		{
			name: "Case 1 sukses handler",
			fields: fields{
				baseURL:  "mock.data",
				endPoint: "testaja",
			},
			args: args{
				IDgarage: 1,
			},
			want: map[int]models.GarageStatus{
				1: {
					OwnerName: "Ceki",
					CarName:   "J",
					IDCar:     "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			R := &RestAPI{
				baseURL:  tt.fields.baseURL,
				endPoint: tt.fields.endPoint,
			}

			srv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					slc := []models.GarageStatus{
						{
							OwnerName: "Ceki",
							CarName:   "J",
							IDCar:     "1",
						},
					}
					dt, _ := json.Marshal(slc)
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, string(dt))
				}))
			R.baseURL = srv.URL

			if got := R.GetGarageStatus(tt.args.IDgarage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestAPI.GetGarageStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
