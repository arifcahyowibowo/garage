package postgres

const (
	// StmtGetGarageByID Statement for get garage by id
	StmtGetGarageByID = `SELECT id_garage, garage_nm, latitude, longtitude from garages where id_garage = $1`
	// StmtGetPositionByGarageID Statement for get postition in garage
	StmtGetPositionByGarageID = `SELECT id_position, id_garage, id_car, position_nm from position where id_garage = $1`
	// StmtGetCarPosition Statement for get Position Car
	StmtGetCarPosition = `select b.garage_nm, b.longtitude, b.latitude, a.position_nm from position a left join garages b on a.id_garage = b.id_garage where a.id_car = $1`
)
