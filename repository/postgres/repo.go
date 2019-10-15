package repository

import (
	"projectapi/models"
)

//go:generate moq -out repo_moq.go . RepoPostgres

// RepoPostgres interface is library for repo
type RepoPostgres interface {
	PrepareQuery()
	GetGaragesByID(garageID int) models.Garages
	GetPositionByGarageID(garageID int) []models.Position
	GetCarPosition(idcar string) []models.CarPosition
}
