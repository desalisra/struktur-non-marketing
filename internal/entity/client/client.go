package client

// City
type City struct {
	CityID   int `db:"Kota_Id" json:"city_id"`
	CityName string `db:"Kota_Name" json:"city_name"`
}