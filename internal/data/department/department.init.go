package department

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
	getDepartments  = "getDepartments"
	qGetDepartments = `SELECT Dpt_Id, CONCAT(Div_Name,' - ', Dpt_Name) Dpt_Name
						FROM M_Divisi
						LEFT JOIN M_Departemen ON Dpt_DivID = Div_ID
						WHERE Div_AktifYN = 'Y' 
						AND Dpt_AktifYN = 'Y'
						ORDER BY Dpt_Name`

	getDepartmentById  = "getDepartmentById"
	qGetDepartmentById = `SELECT Div_Id, Div_Name, Dpt_Id, Dpt_Name
						FROM M_Divisi
						LEFT JOIN M_Departemen ON Dpt_DivID = Div_ID
						WHERE Div_AktifYN = 'Y' 
						AND Dpt_AktifYN = 'Y'
						AND Dpt_Id = ?
						ORDER BY Dpt_Name`		

	getPosition = "getPosition" 	
	qGetPosition = `SELECT Jab_Id, Jab_Jabatan1 
					FROM M_Jabatan
					WHERE Jab_AktifYN = 'Y'
					AND (
						(Jab_DivId = ? AND Jab_DeptId = ?)
						OR (Jab_DivId = ? AND Jab_DeptId IS NULL)
						OR (Jab_DivId IS NULL AND Jab_DeptId = ?)
						OR (Jab_DivId IS NULL AND Jab_DeptId IS NULL)
					)`

)

var (
	readStmt   = []statement{
		{getDepartments, qGetDepartments},
		{getDepartmentById, qGetDepartmentById},
		{getPosition, qGetPosition},
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
