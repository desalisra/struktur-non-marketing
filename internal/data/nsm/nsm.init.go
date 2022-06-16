package nsm

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
	getNsm  = "getNsm"
	qGetNsm = `SELECT Nsm_CompanyId, Pt_Name, Nsm_DepartmentId, Dpt_Name, 
						Nsm_CdGroup, Nsm_Nip, Nsm_Name, Nsm_PositionId, Nsm_Position,
						Nsm_In, Nsm_Out, Nsm_BranchId, Cab_Nama, Nsm_CityId, Kota_Name
					FROM Nm_Rayon_Nsm_202206
					LEFT JOIN M_Pt ON Nsm_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Nsm_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Nsm_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Nsm_CityId = Kota_Id
					WHERE Nsm_ActiveYN = 'Y'
					AND Nsm_CompanyId = ?
					AND Nsm_DepartmentId = ?`
)

var (
	readStmt   = []statement{
		{getNsm, qGetNsm},
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
