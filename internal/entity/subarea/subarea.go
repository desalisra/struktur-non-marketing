package subarea

import (
	"gopkg.in/guregu/null.v3"
)

type ListSubarea struct {
	CompanyId     	int    `db:"Sub_CompanyId" json:"company_id"`
	CompanyName   	string `db:"Pt_Name" json:"company_name"`
	DepartmentID  	int    `db:"Sub_DepartmentId" json:"department_id"`
	DepartmentName	string `db:"Dpt_Name" json:"department_name"`
	CdGrop			string `db:"Sub_CdGroup" json:"code_group"`
	Nip          	string `db:"Sub_Nip" json:"nip"`
	Name         	string `db:"Sub_Name" json:"name"`
	PositionID   	int    `db:"Sub_PositionId" json:"position_id"`
	PositionName 	string `db:"Sub_Position" json:"position_name"`
	DateIn       	null.String   `db:"Sub_In" json:"date_in"`
	DateOut      	null.String   `db:"Sub_Out" json:"date_out"`
	
	BranchID       	int   	`db:"Sub_BranchId" json:"branch_id"`
	BranchName     	string  `db:"Cab_Nama" json:"branch_name"`
	CityID       	int		`db:"Sub_CityId" json:"city_id"`
	CityName      	string  `db:"Kota_Name" json:"city_name"`
}
