package department

// Department
type Department struct {
	DivID   int 	`db:"Div_Id" json:"div_id"`
	DivName string	`db:"Div_Name" json:"div_name"`
	DptID   int		`db:"Dpt_Id" json:"dpt_id"`
	DptName string	`db:"Dpt_Name" json:"dpt_name"`
}

// Position
type Position struct {
	PositionID   int	`db:"Jab_Id" json:"pos_id"`
	PositionName string `db:"Jab_Jabatan1" json:"pos_name"`
}