package area

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
	getArea  = "getArea"
	qGetArea = `SELECT Area_CompanyId, Pt_Name, Area_DepartmentId, Dpt_Name, 
					Area_CdGroup, Area_Nip, Area_Name, Area_PositionId, Area_Position,
					Area_In, Area_Out, Area_DummyYN, Area_BranchId, Cab_Nama, Area_CityId, Kota_Name,
					Area_NipShadow, Area_NameShadow, Area_InShadow, Area_OutShadow, Area_DummyShadowYN,
					Area_Head, Reg_Nip, Reg_Name
				FROM Nm_Rayon_Area_202206
				LEFT JOIN M_Pt ON Area_CompanyId = Pt_Id
				LEFT JOIN M_Departemen ON Area_DepartmentId = Dpt_Id
				LEFT JOIN M_Cabang ON Area_BranchId = Cab_Id
				LEFT JOIN M_Kota ON Area_CityId = Kota_Id
				LEFT JOIN Nm_Rayon_Region_202206 ON Area_Head = Reg_CdGroup
					AND Area_CompanyId = Reg_CompanyId
					AND Area_DepartmentId = Reg_DepartmentId
				WHERE Area_ActiveYN = 'Y'
				AND Area_CompanyId = ?
				AND Area_DepartmentId = ?`
)

var (
	readStmt   = []statement{
		{getArea, qGetArea},
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
