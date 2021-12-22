package entity

type Client struct {
	ID      uint   `gorm:"primarykey"`
	Name    string `gorm:"not null;default:null"`
	HasDebt bool   ``
}

type ClientResultSet struct {
	ResultSet
	Query
	Data []Client
}

func (Client) TableName() string {
	return "Client"
}
