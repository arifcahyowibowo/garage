package db

import (
	md "projectapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetGarageStatus(t *testing.T) {
	type args struct {
		IDGarage int
	}
	tests := []struct {
		name        string
		arguments   args
		wantIsExist md.GarageStatus
	}{
		{
			name: "Test Case 1 garage Exist",
			arguments: args{
				IDGarage: 1,
			},
			wantIsExist: md.GarageStatus{
				OwnerName: "Ceki",
				CarName:   "J",
				IDCar:     "1",
			},
		},
		{
			name: "Test Case 2 garage Not Exist",
			arguments: args{
				IDGarage: 1000000000000,
			},
			wantIsExist: md.GarageStatus{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			gotExist := GetGarageStatus(testcase.arguments.IDGarage)
			assert.Containsf(t, gotExist, testcase.wantIsExist, "Error ")

		})
	}
}
