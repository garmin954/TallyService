package model

// 账本
var LedgerTable string = "ledgers"

type Ledger struct {
	BaseModel
	Name  string `gorm:"type:varchar(50);not null;comment:账本名称" form:"name" json:"name"`
	Desc  string `gorm:"type:tinytext;not null;comment:描述" form:"name" json:"desc"`
	Cover string `gorm:"type:varchar(250);not null;default:'';comment:封面" json:"cover" form:"cover"`
	Uid   int    `gorm:"index;type:int(8);comment:创建人用户ID" json:"-" form:"uid"`
}
