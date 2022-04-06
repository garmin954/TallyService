package model

// 账本
var LedgerTable string = "ledger"

type Ledger struct {
	BaseModel
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(50);not null;comment:账本名称"`
	Desc  string `gorm:"type:tinytext;not null;comment:描述"`
	Cover string `gorm:"type:varchar(250);not null;default:'';comment:封面"`
	Uid   int    `gorm:"index;type:int(8);comment:创建人用户ID"`
}
