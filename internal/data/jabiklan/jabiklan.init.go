package jabiklan

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
	getJabIklan  = "getJabIklan"
	qGetJabIklan = `SELECT A.JabIklan_Id, A.JabIklan_Desc
					FROM M_Jabatan_Iklan A
					LEFT JOIN H_Jabatan_Iklan B ON A.JabIklan_Id = B.JabIklan_Id
					WHERE A.JabIklan_AktifYN = 'Y' AND B.AktifYN = 'Y'
					AND Pt_Id = ?
					AND Jab_Ho = ?
					AND Dept_Ho = ?`
)

var (
	readStmt   = []statement{
		{getJabIklan, qGetJabIklan},
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