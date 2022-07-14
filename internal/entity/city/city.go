package city

// City
type City struct {
	CityID   int    `db:"Kota_Id" json:"city_id"`
	CityName string `db:"Kota_Name" json:"city_name"`
}

// Branch
type Branch struct {
	BranchID      int    `db:"Cab_Id" json:"branch_id"`
	BranchName    string `db:"Cab_Nama" json:"branch_name"`
	BranchAddress string `db:"Cab_Alamat" json:"branch_address"`
}
