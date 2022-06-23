package grpteri

import (
	"gopkg.in/guregu/null.v3"
)

type ResMessage struct {
	Message	string	`json:"msg"`
}

type ListGrpteri struct {
	CompanyId     	int    `db:"Grpt_CompanyId" json:"company_id"`
	CompanyName   	string `db:"Pt_Name" json:"company_name"`
	DepartmentID  	int    `db:"Grpt_DepartmentId" json:"department_id"`
	DepartmentName	string `db:"Dpt_Name" json:"department_name"`
	CdGrop			string `db:"Grpt_CdGroup" json:"code_group"`
	Nip          	string `db:"Grpt_Nip" json:"nip"`
	Name         	string `db:"Grpt_Name" json:"name"`
	PositionID   	int    `db:"Grpt_PositionId" json:"position_id"`
	PositionName 	string 		`db:"Grpt_Position" json:"position_name"`
	DateIn       	null.String `db:"Grpt_In" json:"date_in"`
	DateOut      	null.String `db:"Grpt_Out" json:"date_out"`
	Dummy      		null.String `db:"Grpt_DummyYN" json:"dummy"`
	BranchID       	int   		`db:"Grpt_BranchId" json:"branch_id"`
	BranchName     	string  	`db:"Cab_Nama" json:"branch_name"`
	CityID       	int			`db:"Grpt_CityId" json:"city_id"`
	CityName      	string  	`db:"Kota_Name" json:"city_name"`
	NipShadow       null.String `db:"Grpt_NipShadow" json:"shadow_nip"`
	NameShadow      null.String `db:"Grpt_NameShadow" json:"shadow_name"`
	DateInShadow    null.String  `db:"Grpt_InShadow" json:"shadow_in"`
	DateOutShadow   null.String  `db:"Grpt_OutShadow" json:"shadow_out"`
	DummyShadow	    null.String  `db:"Grpt_DummyShadowYN" json:"shadow_dummy"`
	CdHead	    	null.String  `db:"Grpt_Head" json:"code_head"`
	HeadNip	    	null.String  `db:"Sub_Nip" json:"head_nip"`
	HeadName	    null.String  `db:"Sub_Name" json:"head_name"`
}


type AddGrpteri struct {
	Periode     	string    `json:"periode"`
	CompanyId     	string    `json:"company_id"`
	CompanyName   	string 	  `json:"company_name"`
	DepartmentID  	string    `json:"department_id"`
	DepartmentName	string 	  `json:"department_name"`
	CdGroup			string 	`json:"code_group"`
	Nip          	string `json:"nip"`
	Name         	string `json:"name"`
	PositionID   	string    `json:"position_id"`
	PositionName 	string 		`json:"position_name"`
	DateIn       	null.String `json:"date_in"`
	DateOut      	null.String `json:"date_out"`
	Dummy      		null.String `json:"dummy"`
	BranchID       	string   	`json:"branch_id"`
	BranchName     	string  	`json:"branch_name"`
	CityID       	string		`json:"city_id"`
	CityName      	string  	`json:"city_name"`
	NipShadow       null.String `json:"shadow_nip"`
	NameShadow      null.String `json:"shadow_name"`
	DateInShadow    null.String  `json:"shadow_in"`
	DateOutShadow   null.String  `json:"shadow_out"`
	DummyShadow	    null.String  `json:"shadow_dummy"`
	CdHead	    	null.String  `json:"code_head"`
	HeadNip	    	null.String  `json:"head_nip"`
	HeadName	    null.String  `json:"head_name"`
}