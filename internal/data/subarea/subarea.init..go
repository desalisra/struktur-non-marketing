package subarea

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
	getSubarea  = "getSubarea"
	qGetSubarea = `SELECT Sub_CompanyId, Pt_Name, Sub_DepartmentId, Dpt_Name, 
						Sub_CdGroup, Sub_Nip, Sub_Name, Sub_PositionId, Sub_Position,
						Sub_In, Sub_Out, Sub_DummyYN, Sub_BranchId, Cab_Nama, Sub_CityId, Kota_Name,
						Sub_NipShadow, Sub_NameShadow, Sub_InShadow, Sub_OutShadow, Sub_DummyShadowYN,
						Sub_Head, Area_Nip, Area_Name
					FROM Nm_Rayon_Subarea_202206
					LEFT JOIN M_Pt ON Sub_CompanyId = Pt_Id
					LEFT JOIN M_Departemen ON Sub_DepartmentId = Dpt_Id
					LEFT JOIN M_Cabang ON Sub_BranchId = Cab_Id
					LEFT JOIN M_Kota ON Sub_CityId = Kota_Id
					LEFT JOIN Nm_Rayon_Area_202206 ON Sub_Head = Area_CdGroup
						AND Sub_CompanyId = Area_CompanyId
						AND Sub_DepartmentId = Area_DepartmentId
					WHERE Sub_ActiveYN = 'Y'
					AND Sub_CompanyId = ?
					AND Sub_DepartmentId = ?`
)

var (
	readStmt   = []statement{
		{getSubarea, qGetSubarea},
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
