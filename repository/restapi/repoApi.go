package repository

import (
	"projectapi/models"
)

//go:generate moq -out repoApi_moq.go . RepoCAR

// RepoCAR interface is library for repo
type RepoCAR interface {
	GetGarageStatus(IDgarage int) map[int]models.GarageStatus
}
