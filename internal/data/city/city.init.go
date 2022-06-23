package city

import (
	"context"
	"log"

	"struktur-non-marketing/internal/viper"

	"github.com/jmoiron/sqlx"
)

type (
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)


func (d *Data) UpdateConn() {
	v := viper.GetConn()
	d.db = v.Db
}

// Query List to Prepare
const (
	getCitys  = "getCitys"
	qGetCitys = `SELECT Kota_Id, Kota_Name
				FROM M_Kota
				WHERE Kota_AktifYN = 'Y'
				ORDER BY Kota_Name ASC`
	
	getCityById  = "getCityById"
	qGetCityById = `SELECT Kota_Id, Kota_Name
					FROM M_Kota
					WHERE Kota_AktifYN = 'Y'
					AND Kota_Id = ?
					ORDER BY Kota_Name ASC`
					
	getCityByName  = "getCityByName"
	qGetCityByName = `SELECT Kota_Id, Kota_Name
					FROM M_Kota
					WHERE Kota_AktifYN = 'Y'
					AND Kota_Name LIKE ?
					ORDER BY Kota_Name ASC`

	getBranchByCityId  = "getBranchByCityId"
	qGetBranchByCityId = `SELECT Cab_Id, Cab_Nama, Cab_Alamat 
							FROM M_Cabang
							WHERE Cab_AktifYN = 'Y'
							AND Cab_KotaId = ?
							ORDER BY Cab_Nama ASC`
)

var (
	readStmt   = []statement{
		{getCitys, qGetCitys},
		{getCityById, qGetCityById},
		{getCityByName, qGetCityByName},
		{getBranchByCityId, qGetBranchByCityId},
	}
	upsertStmt = []statement{}
	deleteStmt = []statement{}
)

func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range upsertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize upsert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}
