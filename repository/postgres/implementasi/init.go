package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Repo is object for database
type Repo struct {
	DbObject                  *sql.DB
	prepGetGarageByID         *sql.Stmt
	prepGetPositionByGarageID *sql.Stmt
	prepGetCarPosition        *sql.Stmt
}

// New it Open connection to server
func New(db *sql.DB) (*Repo, error) {
	return &Repo{
		DbObject: db,
	}, nil
}

// PrepareQuery is function to init query
func (repo *Repo) PrepareQuery() {
	var err error
	repo.prepGetGarageByID, err = repo.DbObject.Prepare(StmtGetGarageByID)
	if err != nil {
		log.Println("Error prepare StmtGetGarageByID")
	}
	repo.prepGetPositionByGarageID, err = repo.DbObject.Prepare(StmtGetPositionByGarageID)
	if err != nil {
		log.Println("Error prepare StmtGetPositionByGarageID")
	}
	repo.prepGetCarPosition, err = repo.DbObject.Prepare(StmtGetCarPosition)
	if err != nil {
		log.Println("Error prepare StmtGetCarPosition")
	}
}
