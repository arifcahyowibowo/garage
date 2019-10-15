package usecase

import (
	postgresrepo "projectapi/repository/postgres"
	postgres "projectapi/repository/postgres/implementasi"
	restapirepo "projectapi/repository/restapi"
	restapi "projectapi/repository/restapi/implementasi"
)

// GarageServiceStruct is struct as Object in Service
type GarageServiceStruct struct {
	Postgres postgresrepo.RepoPostgres
	Cardata  restapirepo.RepoCAR
}

// New is constructor for GarageService
func New(postgresRepo *postgres.Repo, carData *restapi.RestAPI) (Service *GarageServiceStruct) {
	Service = &GarageServiceStruct{
		Postgres: postgresRepo,
		Cardata:  carData,
	}
	return Service
}
