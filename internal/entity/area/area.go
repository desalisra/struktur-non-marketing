package area

import (
	"gopkg.in/guregu/null.v3"
)

type ListArea struct {
	CompanyId     	int    `db:"Area_CompanyId" json:"company_id"`
	CompanyName   	string `db:"Pt_Name" json:"company_name"`
	DepartmentID  	int    `db:"Area_DepartmentId" json:"department_id"`
	DepartmentName	string `db:"Dpt_Name" json:"department_name"`
	CdGrop			string `db:"Area_CdGroup" json:"code_group"`
	Nip          	string `db:"Area_Nip" json:"nip"`
	Name         	string `db:"Area_Name" json:"name"`
	PositionID   	int    `db:"Area_PositionId" json:"position_id"`
	PositionName 	string `db:"Area_Position" json:"position_name"`
	DateIn       	null.String   `db:"Area_In" json:"date_in"`
	DateOut      	null.String   `db:"Area_Out" json:"date_out"`
	
	BranchID       	int   	`db:"Area_BranchId" json:"branch_id"`
	BranchName     	string  `db:"Cab_Nama" json:"branch_name"`
	CityID       	int		`db:"Area_CityId" json:"city_id"`
	CityName      	string  `db:"Kota_Name" json:"city_name"`
}
