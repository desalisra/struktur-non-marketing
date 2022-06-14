package city

// City
type City struct {
	CityID   int `db:"Kota_Id" json:"city_id"`
	CityName string `db:"Kota_Name" json:"city_name"`
}

// Bracnh
type Bracnh struct {
	BracnhID   int `db:"Cab_Id" json:"branch_id"`
	BracnhName string `db:"Cab_Nama" json:"branch_name"`
	BracnhAddress string `db:"Cab_Alamat" json:"branch_address"`
}