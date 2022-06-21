package karyawan

import "gopkg.in/guregu/null.v3"

// ListKaryawan
type ListKaryawan struct {
	Nip			null.String 	`db:"Kry_Nip" json:"nip"`
	Nama		null.String	`db:"Kry_Nama" json:"name"`
	Status  	null.String	`db:"Kry_StatusKry" json:"status"`
	TglMasuk	null.String	`db:"Kry_TglMasukAsli" json:"tgl_masuk"`
	JabId 		null.String	`db:"Kry_JabId" json:"jab_id"`
	JabName 	null.String	`db:"Kry_Jabatan" json:"jab_name"`
}
