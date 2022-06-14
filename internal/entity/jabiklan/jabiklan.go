package jabiklan

type JabIklan struct {
	IklanID   int    `db:"JabIklan_Id" json:"iklan_id"`
	IklanName string `db:"JabIklan_Desc" json:"iklan_name"`
}