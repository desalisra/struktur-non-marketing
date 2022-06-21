package nsm

import (
	"gopkg.in/guregu/null.v3"
)

type ListNsm struct {
	CompanyId     	int    `db:"Nsm_CompanyId" json:"company_id"`
	CompanyName   	string `db:"Pt_Name" json:"company_name"`
	DepartmentID  	int    `db:"Nsm_DepartmentId" json:"department_id"`
	DepartmentName	string `db:"Dpt_Name" json:"department_name"`
	CdGrop			string `db:"Nsm_CdGroup" json:"code_group"`
	Nip          	string `db:"Nsm_Nip" json:"nip"`
	Name         	string `db:"Nsm_Name" json:"name"`
	PositionID   	int    `db:"Nsm_PositionId" json:"position_id"`
	PositionName 	string `db:"Nsm_Position" json:"position_name"`
	DateIn       	null.String   `db:"Nsm_In" json:"date_in"`
	DateOut      	null.String   `db:"Nsm_Out" json:"date_out"`
	Dummy      		null.String   `db:"Nsm_DummyYN" json:"dummy"`
	BranchID       	int   	`db:"Nsm_BranchId" json:"branch_id"`
	BranchName     	string  `db:"Cab_Nama" json:"branch_name"`
	CityID       	int		`db:"Nsm_CityId" json:"city_id"`
	CityName      	string  `db:"Kota_Name" json:"city_name"`
	NipShadow       null.String `db:"Nsm_NipShadow" json:"shadow_nip"`
	NameShadow      null.String `db:"Nsm_NameShadow" json:"shadow_name"`
	DateInShadow    null.String   `db:"Nsm_InShadow" json:"shadow_in"`
	DateOutShadow   null.String   `db:"Nsm_OutShadow" json:"shadow_out"`
	DummyShadow	    null.String   `db:"Nsm_DummyShadowYN" json:"shadow_dummy"`
}
