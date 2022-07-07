package grpteri

type ResMessage struct {
	Message string  `json:"message"`
	Data    Grpteri `json:"data"`
}

type Grpteri struct {
	Periode        string `db:"Grpt_Periode" json:"periode"`
	CompanyId      string `db:"Grpt_CompanyId" json:"company_id"`
	CompanyName    string `db:"Pt_Name" json:"company_name"`
	DepartmentId   string `db:"Grpt_DepartmentId" json:"department_id"`
	DepartmentName string `db:"Dpt_Name" json:"department_name"`
	CdGroup        string `db:"Grpt_CdGroup" json:"code_group"`
	Nip            string `db:"Grpt_Nip" json:"nip"`
	Name           string `db:"Grpt_Name" json:"name"`
	PositionId     string `db:"Grpt_PositionId" json:"position_id"`
	PositionName   string `db:"Grpt_Position" json:"position_name"`
	DateIn         string `db:"Grpt_In" json:"date_in"`
	DateOut        string `db:"Grpt_Out" json:"date_out"`
	Dummy          string `db:"Grpt_DummyYN" json:"dummy"`
	BranchId       string `db:"Grpt_BranchId" json:"branch_id"`
	BranchName     string `db:"Cab_Nama" json:"branch_name"`
	CityId         string `db:"Grpt_CityId" json:"city_id"`
	CityName       string `db:"Kota_Name" json:"city_name"`
	NipShadow      string `db:"Grpt_NipShadow" json:"shadow_nip"`
	NameShadow     string `db:"Grpt_NameShadow" json:"shadow_name"`
	DateInShadow   string `db:"Grpt_InShadow" json:"shadow_in"`
	DateOutShadow  string `db:"Grpt_OutShadow" json:"shadow_out"`
	DummyShadow    string `db:"Grpt_DummyShadowYN" json:"shadow_dummy"`
	CdHead         string `db:"Grpt_Head" json:"code_head"`
	HeadNip        string `db:"Sub_Nip" json:"head_nip"`
	HeadName       string `db:"Sub_Name" json:"head_name"`
}
