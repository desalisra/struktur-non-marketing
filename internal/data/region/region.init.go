package region

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
	getRegion  = "getRegion"
	qGetRegion = `SELECT Reg_CompanyId, Pt_Name, Reg_DepartmentId, Dpt_Name, 
						Reg_CdGroup, Reg_Nip, Reg_Name, Reg_PositionId, Reg_Position,
						Reg_In, Reg_Out, Reg_BranchId, Cab_Nama, Reg_CityId, Kota_Name
					FROM Nm_Rayon_Region_202206
					LEFT JOIN M_Pt ON Reg_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Reg_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Reg_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Reg_CityId = Kota_Id
					WHERE Reg_ActiveYN = 'Y'
					AND Reg_CompanyId = ?
					AND Reg_DepartmentId = ?`
)

var (
	readStmt   = []statement{
		{getRegion, qGetRegion},
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
