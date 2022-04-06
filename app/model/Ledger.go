package model

var LedgerTable string = "ledger"

type Ledger struct {
	BaseModel
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"varchat(50)";not null`
	UserId User   `gorm:"foreignKey:User.ID"`
}
