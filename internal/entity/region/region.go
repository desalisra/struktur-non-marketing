package region

import (
	"gopkg.in/guregu/null.v3"
)

type ListRegion struct {
	CompanyId     	int    `db:"Reg_CompanyId" json:"company_id"`
	CompanyName   	string `db:"Pt_Name" json:"company_name"`
	DepartmentID  	int    `db:"Reg_DepartmentId" json:"department_id"`
	DepartmentName	string `db:"Dpt_Name" json:"department_name"`
	CdGrop			string `db:"Reg_CdGroup" json:"code_group"`
	Nip          	string `db:"Reg_Nip" json:"nip"`
	Name         	string `db:"Reg_Name" json:"name"`
	PositionID   	int    `db:"Reg_PositionId" json:"position_id"`
	PositionName 	string `db:"Reg_Position" json:"position_name"`
	DateIn       	null.String   `db:"Reg_In" json:"date_in"`
	DateOut      	null.String   `db:"Reg_Out" json:"date_out"`
	Dummy      		null.String   `db:"Reg_DummyYN" json:"dummy"`
	BranchID       	int   	`db:"Reg_BranchId" json:"branch_id"`
	BranchName     	string  `db:"Cab_Nama" json:"branch_name"`
	CityID       	int		`db:"Reg_CityId" json:"city_id"`
	CityName      	string  `db:"Kota_Name" json:"city_name"`
	NipShadow       null.String `db:"Reg_NipShadow" json:"shadow_nip"`
	NameShadow      null.String `db:"Reg_NameShadow" json:"shadow_name"`
	DateInShadow    null.String   `db:"Reg_InShadow" json:"shadow_in"`
	DateOutShadow   null.String   `db:"Reg_OutShadow" json:"shadow_out"`
	DummyShadow	    null.String   `db:"Reg_DummyShadowYN" json:"shadow_dummy"`
	CdHead	    	null.String   		  `db:"Reg_Head" json:"code_head"`
	HeadNip	    	null.String   `db:"Nsm_Nip" json:"head_nip"`
	HeadName	    null.String   `db:"Nsm_Name" json:"head_name"`
}
