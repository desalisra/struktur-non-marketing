package grpteri

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
	getGrpteri  = "getGrpteri"
	qGetGrpteri = `SELECT Grpt_CompanyId, Pt_Name, Grpt_DepartmentId, Dpt_Name, 
						Grpt_CdGroup, Grpt_Nip, Grpt_Name, Grpt_PositionId, Grpt_Position,
						Grpt_In, Grpt_Out, Grpt_DummyYN, Grpt_BranchId, Cab_Nama, Grpt_CityId, Kota_Name,
						Grpt_NipShadow, Grpt_NameShadow, Grpt_InShadow, Grpt_OutShadow, Grpt_DummyShadowYN,
						Grpt_Head, Sub_Nip, Sub_Name
					FROM Nm_Rayon_Grpteri_202206
					LEFT JOIN M_Pt ON Grpt_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Grpt_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Grpt_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Grpt_CityId = Kota_Id
					LEFT JOIN Nm_Rayon_Subarea_202206 ON Grpt_Head = Sub_CdGroup
						AND Grpt_CompanyId = Sub_CompanyId
						AND Grpt_DepartmentId = Sub_DepartmentId
					WHERE Grpt_ActiveYN = 'Y'
					AND Grpt_CompanyId = ?
					AND Grpt_DepartmentId = ?
					`
)

var (
	readStmt   = []statement{
		{getGrpteri, qGetGrpteri},
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
