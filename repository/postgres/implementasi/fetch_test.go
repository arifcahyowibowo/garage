package postgres

import (
	"database/sql"
	"projectapi/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetGaragesByID(t *testing.T) {
	type fields struct {
		DbObject                  *sql.DB
		prepGetGarageByID         *sql.Stmt
		prepGetPositionByGarageID *sql.Stmt
		prepGetCarPosition        *sql.Stmt
	}
	type args struct {
		garageID int
	}
	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	tests := []struct {
		name   string
		fields fields
		rows   []string
		args   args
		want   models.Garages
	}{
		{
			name: "Testcase 1 with returned Data",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				garageID: 1,
			},
			rows: []string{
				"1",
				"Garage 1",
				"0",
				"0",
			},
			want: models.Garages{
				IDGarage:   1,
				GarageNm:   "Garage 1",
				Latitude:   "0",
				Longtitude: "0",
			},
		},
		{
			name: "Testcase 2 with 0 id",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				garageID: 0,
			},
			want: models.Garages{
				IDGarage: 0,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{
				DbObject:                  tt.fields.DbObject,
				prepGetGarageByID:         tt.fields.prepGetGarageByID,
				prepGetPositionByGarageID: tt.fields.prepGetPositionByGarageID,
				prepGetCarPosition:        tt.fields.prepGetCarPosition,
			}

			rows := sqlmock.NewRows([]string{
				"id_garage",
				"garage_nm",
				"latitude",
				"longtitude",
			})

			if tt.rows != nil {
				rows.AddRow(tt.rows[0], tt.rows[1], tt.rows[2], tt.rows[2])
			}

			mockSQL.ExpectPrepare("^SELECT (.*)").
				ExpectQuery().
				WithArgs(tt.args.garageID).
				WillReturnRows(rows)

			repo.prepGetGarageByID, _ = mockDB.Prepare(StmtGetGarageByID)
			got := repo.GetGaragesByID(tt.args.garageID)
			assert.Equalf(t, tt.want, got, "Kenapa Error")

		})
	}
}

func TestRepo_GetPositionByGarageID(t *testing.T) {
	type fields struct {
		DbObject                  *sql.DB
		prepGetGarageByID         *sql.Stmt
		prepGetPositionByGarageID *sql.Stmt
		prepGetCarPosition        *sql.Stmt
	}
	type args struct {
		garageID int
	}

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	tests := []struct {
		name   string
		fields fields
		rows   []string
		args   args
		want   []models.Position
	}{
		{
			name: "Testcase 1 with returned Data",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				garageID: 1,
			},
			rows: []string{
				"1",
				"1",
				"1",
				"Slot 1",
			},
			want: []models.Position{
				models.Position{
					IDPosition:   1,
					IDGarage:     1,
					IDCar:        1,
					PositionName: "Slot 1",
				},
			},
		},
		{
			name: "Testcase 2 with 0 id",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				garageID: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{
				DbObject:                  tt.fields.DbObject,
				prepGetGarageByID:         tt.fields.prepGetGarageByID,
				prepGetPositionByGarageID: tt.fields.prepGetPositionByGarageID,
				prepGetCarPosition:        tt.fields.prepGetCarPosition,
			}

			rows := sqlmock.NewRows([]string{
				"IDPosition",
				"IDGarage",
				"IDCar",
				"PositionName",
			})

			if tt.rows != nil {
				rows.AddRow(tt.rows[0], tt.rows[1], tt.rows[2], tt.rows[3])
			}

			mockSQL.ExpectPrepare("^SELECT (.+)").
				ExpectQuery().
				WithArgs(tt.args.garageID).
				WillReturnRows(rows)

			repo.prepGetPositionByGarageID, _ = mockDB.Prepare(StmtGetPositionByGarageID)

			got := repo.GetPositionByGarageID(tt.args.garageID)
			assert.Equalf(t, tt.want, got, "Kenapa Error")

		})
	}
}

func TestRepo_GetCarPosition(t *testing.T) {
	type fields struct {
		DbObject                  *sql.DB
		prepGetGarageByID         *sql.Stmt
		prepGetPositionByGarageID *sql.Stmt
		prepGetCarPosition        *sql.Stmt
	}
	type args struct {
		idcar string
	}

	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()
	tests := []struct {
		name   string
		fields fields
		args   args
		rows   []string
		want   []models.CarPosition
	}{
		{
			name: "Testcase 1 with returned Data",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				idcar: "1",
			},
			rows: []string{
				"E4 Garage",
				"0",
				"0",
				"Slot 2",
			},
			want: []models.CarPosition{
				models.CarPosition{
					GarageName:   "E4 Garage",
					Longtitude:   "0",
					Latitude:     "0",
					PositionName: "Slot 2",
				},
			},
		},
		{
			name: "Testcase 2 with 0 id",
			fields: fields{
				DbObject: mockDB,
			},
			args: args{
				idcar: "",
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{
				DbObject:                  tt.fields.DbObject,
				prepGetGarageByID:         tt.fields.prepGetGarageByID,
				prepGetPositionByGarageID: tt.fields.prepGetPositionByGarageID,
				prepGetCarPosition:        tt.fields.prepGetCarPosition,
			}

			rows := sqlmock.NewRows([]string{
				"IDPosition",
				"IDGarage",
				"IDCar",
				"PositionName",
			})

			if tt.rows != nil {
				rows.AddRow(tt.rows[0], tt.rows[1], tt.rows[2], tt.rows[3])
			}

			mockSQL.ExpectPrepare("^SELECT (.+)").
				ExpectQuery().
				WithArgs(tt.args.idcar).
				WillReturnRows(rows)

			repo.prepGetCarPosition, _ = mockDB.Prepare(StmtGetPositionByGarageID)

			got := repo.GetCarPosition(tt.args.idcar)
			assert.Equalf(t, tt.want, got, "Kenapa Error")
		})
	}
}
